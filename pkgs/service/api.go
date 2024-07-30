package service

import (
	"Listen/config"
	"Listen/pkgs/prost"
	"Listen/pkgs/redis"
	"context"
	"encoding/json"
	"fmt"
	"github.com/cenkalti/backoff/v4"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"math/big"
	"net/http"
	"strconv"
	"sync"
)

var (
	rewardBasePoints   *big.Int
	dailySnapshotQuota *big.Int
)

const baseExponent = int64(100000000)

type DailyRewardsRequest struct {
	SlotID int    `json:"slot_id"`
	Day    int    `json:"day"`
	Token  string `json:"token"`
}

type RewardsRequest struct {
	SlotID int    `json:"slot_id"`
	Token  string `json:"token"`
}

type SubmissionsRequest struct {
	SlotID   int    `json:"slot_id"`
	Token    string `json:"token"`
	PastDays int    `json:"past_days"`
}

type DailySubmissions struct {
	Day         int   `json:"day"`
	Submissions int64 `json:"submissions"`
}

func getTotalRewards(slotId int) (int, error) {
	ctx := context.Background()
	slotIdStr := strconv.Itoa(slotId)

	totalSlotRewardsKey := redis.TotalSlotRewards(slotIdStr)
	totalSlotRewards, err := redis.Get(ctx, totalSlotRewardsKey)

	if err != nil || totalSlotRewards == "" {
		slotIdBigInt := big.NewInt(int64(slotId))
		totalSlotRewardsBigInt, err := FetchSlotRewardsPoints(slotIdBigInt)
		if err != nil {
			return 0, err
		}

		totalSlotRewardsInt := totalSlotRewardsBigInt.Int64()
		// TODO: Try to avoid query
		day, err := prost.MustQuery[*big.Int](context.Background(), prost.Instance.DayCounter)
		if err != nil {
			log.Errorln("Unable to query day from contract")
		}
		currentDayRewardsKey := redis.SlotRewardsForDay(slotIdStr, day.String())
		currentDayRewards, err := redis.Get(ctx, currentDayRewardsKey)

		if err != nil || currentDayRewards == "" {
			currentDayRewards = "0"
		}

		currentDayRewardsInt, err := strconv.ParseInt(currentDayRewards, 10, 64)

		totalSlotRewardsInt += currentDayRewardsInt

		redis.Set(ctx, totalSlotRewardsKey, totalSlotRewards, 0)

		return int(totalSlotRewardsInt), nil
	}

	parsedRewards, err := strconv.ParseInt(totalSlotRewards, 10, 64)
	if err != nil {
		return 0, err
	}

	return int(parsedRewards), nil
}

func FetchSlotRewardsPoints(slotId *big.Int) (*big.Int, error) {
	var err error

	slotIdStr := slotId.String()
	var points *big.Int

	retryErr := backoff.Retry(func() error {
		points, err = prost.Instance.SlotRewardPoints(&bind.CallOpts{}, slotId)
		return err
	}, backoff.WithMaxRetries(backoff.NewExponentialBackOff(), 5))

	if retryErr != nil {
		log.Errorf("Unable to query SlotRewardPoints for slot %s: %v", slotIdStr, retryErr)
		return nil, retryErr
	}

	return points, nil
}

func getDailySubmissions(slotId int, day *big.Int) int64 {
	if val, err := redis.Get(context.Background(), redis.SlotSubmissionKey(strconv.Itoa(slotId), day.String())); err != nil || val == "" {
		subs, err := prost.MustQuery[*big.Int](context.Background(), func(opts *bind.CallOpts) (*big.Int, error) {
			subs, err := prost.Instance.SlotSubmissionCount(opts, big.NewInt(int64(slotId)), day)
			return subs, err
		})
		if err != nil {
			log.Errorln("Could not fetch submissions from contract: ", err.Error())
			return 0
		}
		return subs.Int64()
	} else {
		submissions, _ := new(big.Int).SetString(val, 10)
		return submissions.Int64()
	}
}

func handleTotalSubmissions(w http.ResponseWriter, r *http.Request) {
	var request SubmissionsRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Authenticate token
	if request.Token != config.SettingsObj.AuthReadToken {
		http.Error(w, "Incorrect Token!", http.StatusUnauthorized)
		return
	}

	if request.PastDays < 1 {
		http.Error(w, "Past days should be at least 1", http.StatusBadRequest)
		return
	}

	slotID := request.SlotID
	if slotID < 1 || slotID > 10000 {
		http.Error(w, fmt.Sprintf("Invalid slotId: %d", slotID), http.StatusBadRequest)
		return
	}

	// TODO: Try to avoid query
	day, err := prost.MustQuery[*big.Int](context.Background(), prost.Instance.DayCounter)
	if err != nil {
		log.Errorln("Unable to query day from contract")
	}

	currentDay := new(big.Int).Set(day)
	submissionsResponse := make([]DailySubmissions, request.PastDays)

	var wg sync.WaitGroup
	ch := make(chan DailySubmissions, request.PastDays)

	for i := 0; i < request.PastDays; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			day := new(big.Int).Sub(currentDay, big.NewInt(int64(i)))
			subs := getDailySubmissions(request.SlotID, day)
			ch <- DailySubmissions{Day: int(day.Int64()), Submissions: subs}
		}(i)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for submission := range ch {
		submissionsResponse[int(currentDay.Int64())-submission.Day] = submission
	}

	response := struct {
		Info struct {
			Success  bool               `json:"success"`
			Response []DailySubmissions `json:"response"`
		} `json:"info"`
		RequestID string `json:"request_id"`
	}{
		Info: struct {
			Success  bool               `json:"success"`
			Response []DailySubmissions `json:"response"`
		}{
			Success:  true,
			Response: submissionsResponse,
		},
		RequestID: r.Context().Value("request_id").(string),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// FetchContractConstants fetches constants from the contract
func FetchContractConstants(ctx context.Context) error {
	var err error

	dailySnapshotQuota, err = prost.MustQuery[*big.Int](ctx, prost.Instance.DailySnapshotQuota)

	if err != nil {
		return fmt.Errorf("failed to fetch contract DailySnapshotQuota: %v", err)
	}

	rewardBasePoints, err = prost.MustQuery[*big.Int](ctx, prost.Instance.RewardBasePoints)

	if err != nil {
		return fmt.Errorf("failed to fetch contract RewardBasePoints: %v", err)
	}

	return nil
}

func FetchSlotSubmissionCount(slotID *big.Int, day *big.Int) (*big.Int, error) {
	var count *big.Int
	var err error

	retryErr := backoff.Retry(func() error {
		count, err = prost.Instance.SlotSubmissionCount(&bind.CallOpts{}, slotID, day)
		return err
	}, backoff.WithMaxRetries(backoff.NewExponentialBackOff(), 5))

	if retryErr != nil {
		return nil, fmt.Errorf("failed to fetch submission count: %v", retryErr)
	}

	return count, nil
}

func handleDailyRewards(w http.ResponseWriter, r *http.Request) {
	var dailySubmissionCount int64

	var request DailyRewardsRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Authenticate token
	if request.Token != config.SettingsObj.AuthReadToken {
		http.Error(w, "Incorrect Token!", http.StatusUnauthorized)
		return
	}

	slotID := request.SlotID
	if slotID < 1 || slotID > 10000 {
		http.Error(w, fmt.Sprintf("Invalid slotId: %d", slotID), http.StatusBadRequest)
		return
	}

	day := int64(request.Day)

	var rewards int
	key := redis.SlotSubmissionKey(strconv.Itoa(slotID), strconv.Itoa(request.Day))

	// Try to get from Redis
	dailySubmissionCountFromCache, err := redis.Get(r.Context(), key)

	if err != nil || dailySubmissionCountFromCache == "" {
		slotIDBigInt := big.NewInt(int64(slotID))
		count, err := FetchSlotSubmissionCount(slotIDBigInt, big.NewInt(day))
		if err != nil {
			http.Error(w, "Failed to fetch submission count: "+err.Error(), http.StatusInternalServerError)
			return
		}

		dailySubmissionCount = count.Int64()
		slotRewardsForDayKey := redis.SlotRewardsForDay(strconv.Itoa(slotID), strconv.Itoa(request.Day))
		redis.Set(r.Context(), slotRewardsForDayKey, strconv.FormatInt(dailySubmissionCount, 10), 0)
	} else if dailySubmissionCountFromCache != "" {
		dailyCount, err := strconv.ParseInt(dailySubmissionCountFromCache, 10, 64)
		if err != nil {
			log.Errorf("Failed to parse daily submission count from cache: %v", err)
		}
		dailySubmissionCount = dailyCount
	}

	log.Debugln("DailySnapshotQuota: ", dailySnapshotQuota)

	cmp := big.NewInt(dailySubmissionCount).Cmp(dailySnapshotQuota)
	if cmp >= 0 {
		rewards = int(new(big.Int).Div(rewardBasePoints, big.NewInt(baseExponent)).Int64())
	} else {
		rewards = 0
	}

	response := struct {
		Info struct {
			Success  bool `json:"success"`
			Response int  `json:"response"`
		} `json:"info"`
		RequestID string `json:"request_id"`
	}{
		Info: struct {
			Success  bool `json:"success"`
			Response int  `json:"response"`
		}{
			Success:  true,
			Response: rewards,
		},
		RequestID: r.Context().Value("request_id").(string),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func handleTotalRewards(w http.ResponseWriter, r *http.Request) {
	var request RewardsRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Authenticate token
	if request.Token != config.SettingsObj.AuthReadToken {
		http.Error(w, "Incorrect Token!", http.StatusUnauthorized)
		return
	}

	slotID := request.SlotID
	if slotID < 1 || slotID > 10000 {
		http.Error(w, fmt.Sprintf("Invalid slotId: %d", slotID), http.StatusBadRequest)
		return
	}

	slotRewardPoints, err := getTotalRewards(slotID)
	if err != nil {
		http.Error(w, "Failed to fetch slot reward points: "+err.Error(), http.StatusInternalServerError)
		return
	}

	response := struct {
		Info struct {
			Success  bool `json:"success"`
			Response int  `json:"response"`
		} `json:"info"`
		RequestID string `json:"request_id"`
	}{
		Info: struct {
			Success  bool `json:"success"`
			Response int  `json:"response"`
		}{
			Success:  true,
			Response: slotRewardPoints,
		},
		RequestID: r.Context().Value("request_id").(string),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func RequestMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestID := uuid.New().String()
		ctx := context.WithValue(r.Context(), "request_id", requestID)
		r = r.WithContext(ctx)

		log.WithField("request_id", requestID).Infof("Request started for: %s", r.URL.Path)

		w.Header().Set("X-Request-ID", requestID)

		next.ServeHTTP(w, r)

		log.WithField("request_id", requestID).Infof("Request ended")
	})
}

func StartApiServer() {
	err := FetchContractConstants(context.Background())
	if err != nil {
		log.Errorf("Failed to fetch contract constants: %v", err)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/getTotalRewards", handleTotalRewards)
	mux.HandleFunc("/getDailyRewards", handleDailyRewards)
	mux.HandleFunc("/totalSubmissions", handleTotalSubmissions)

	handler := RequestMiddleware(mux)

	log.Println("Server is running on port 9988")
	log.Fatal(http.ListenAndServe(":9988", handler))
}
