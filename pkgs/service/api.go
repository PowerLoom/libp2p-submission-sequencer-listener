package service

import (
	"Listen/config"
	"Listen/pkgs"
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
	"strings"
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

type PastEpochsRequest struct {
	Token      string `json:"token"`
	PastEpochs int    `json:"past_epochs"`
}

type PastBatchesRequest struct {
	Token       string `json:"token"`
	PastBatches int    `json:"past_batches"`
}

type DailySubmissions struct {
	Day         int   `json:"day"`
	Submissions int64 `json:"submissions"`
}

func getTotalRewards(slotId int) (int64, error) {
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

		var day *big.Int
		dayStr, _ := redis.Get(context.Background(), pkgs.SequencerDayKey)
		if dayStr == "" {
			//	TODO: Report unhandled error
			day, _ = prost.MustQuery[*big.Int](context.Background(), prost.Instance.DayCounter)
		} else {
			day, _ = new(big.Int).SetString(dayStr, 10)
		}

		currentDayRewardsKey := redis.SlotRewardsForDay(slotIdStr, day.String())
		currentDayRewards, err := redis.Get(ctx, currentDayRewardsKey)

		if err != nil || currentDayRewards == "" {
			currentDayRewards = "0"
		}

		currentDayRewardsInt, err := strconv.ParseInt(currentDayRewards, 10, 64)

		totalSlotRewardsInt += currentDayRewardsInt

		redis.Set(ctx, totalSlotRewardsKey, totalSlotRewards, 0)

		return totalSlotRewardsInt, nil
	}

	parsedRewards, err := strconv.ParseInt(totalSlotRewards, 10, 64)
	if err != nil {
		return 0, err
	}

	return parsedRewards, nil
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

	var day *big.Int
	dayStr, _ := redis.Get(context.Background(), pkgs.SequencerDayKey)
	if dayStr == "" {
		//	TODO: Report unhandled error
		day, _ = prost.MustQuery[*big.Int](context.Background(), prost.Instance.DayCounter)
	} else {
		day, _ = new(big.Int).SetString(dayStr, 10)
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

func handleFinalizedBatchSubmissions(w http.ResponseWriter, r *http.Request) {
	var request PastEpochsRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Authenticate token
	if request.Token != config.SettingsObj.AuthReadToken {
		http.Error(w, "Incorrect Token!", http.StatusUnauthorized)
		return
	}

	pastEpochs := request.PastEpochs
	if pastEpochs < 0 {
		http.Error(w, "Past epochs should be at least 0", http.StatusBadRequest)
		return
	}

	keys := redis.RedisClient.Keys(context.Background(), fmt.Sprintf("%s.finalized_batch_submissions.*", pkgs.ProcessTriggerKey)).Val()

	var logs []map[string]interface{}

	for _, key := range keys {
		entry, err := redis.Get(context.Background(), key)
		if err != nil {
			continue
		}

		var logEntry map[string]interface{}
		err = json.Unmarshal([]byte(entry), &logEntry)
		if err != nil {
			continue
		}

		logs = append(logs, logEntry)
	}

	if pastEpochs > 0 && len(logs) > pastEpochs {
		logs = logs[len(logs)-pastEpochs:]
	}

	response := struct {
		Info struct {
			Success bool                     `json:"success"`
			Logs    []map[string]interface{} `json:"logs"`
		} `json:"info"`
		RequestID string `json:"request_id"`
	}{
		Info: struct {
			Success bool                     `json:"success"`
			Logs    []map[string]interface{} `json:"logs"`
		}{
			Success: true,
			Logs:    logs,
		},
		RequestID: r.Context().Value("request_id").(string),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func handleTriggeredCollectionFlows(w http.ResponseWriter, r *http.Request) {
	var request PastEpochsRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Authenticate token
	if request.Token != config.SettingsObj.AuthReadToken {
		http.Error(w, "Incorrect Token!", http.StatusUnauthorized)
		return
	}

	pastEpochs := request.PastEpochs
	if pastEpochs < 0 {
		http.Error(w, "Past epochs should be at least 0", http.StatusBadRequest)
		return
	}

	keys := redis.RedisClient.Keys(context.Background(), fmt.Sprintf("%s.triggered_collection_flow.*", pkgs.ProcessTriggerKey)).Val()

	var logs []map[string]interface{}

	for _, key := range keys {
		entry, err := redis.Get(context.Background(), key)
		if err != nil {
			continue
		}

		var logEntry map[string]interface{}
		err = json.Unmarshal([]byte(entry), &logEntry)
		if err != nil {
			continue
		}

		logs = append(logs, logEntry)
	}

	if pastEpochs > 0 && len(logs) > pastEpochs {
		logs = logs[len(logs)-pastEpochs:]
	}

	response := struct {
		Info struct {
			Success bool                     `json:"success"`
			Logs    []map[string]interface{} `json:"logs"`
		} `json:"info"`
		RequestID string `json:"request_id"`
	}{
		Info: struct {
			Success bool                     `json:"success"`
			Logs    []map[string]interface{} `json:"logs"`
		}{
			Success: true,
			Logs:    logs,
		},
		RequestID: r.Context().Value("request_id").(string),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func handleBuiltBatches(w http.ResponseWriter, r *http.Request) {
	var request PastBatchesRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Authenticate token
	if request.Token != config.SettingsObj.AuthReadToken {
		http.Error(w, "Incorrect Token!", http.StatusUnauthorized)
		return
	}

	pastBatches := request.PastBatches
	if pastBatches < 0 {
		http.Error(w, "Past batches should be at least 0", http.StatusBadRequest)
		return
	}

	keys := redis.RedisClient.Keys(context.Background(), fmt.Sprintf("%s.build_batch.*", pkgs.ProcessTriggerKey)).Val()

	var logs []map[string]interface{}

	for _, key := range keys {
		entry, err := redis.Get(context.Background(), key)
		if err != nil {
			continue
		}

		var logEntry map[string]interface{}
		err = json.Unmarshal([]byte(entry), &logEntry)
		if err != nil {
			continue
		}

		logs = append(logs, logEntry)
	}

	if pastBatches > 0 && len(logs) > pastBatches {
		logs = logs[len(logs)-pastBatches:]
	}

	response := struct {
		Info struct {
			Success bool                     `json:"success"`
			Logs    []map[string]interface{} `json:"logs"`
		} `json:"info"`
		RequestID string `json:"request_id"`
	}{
		Info: struct {
			Success bool                     `json:"success"`
			Logs    []map[string]interface{} `json:"logs"`
		}{
			Success: true,
			Logs:    logs,
		},
		RequestID: r.Context().Value("request_id").(string),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func handleCommittedSubmissionBatches(w http.ResponseWriter, r *http.Request) {
	var request PastBatchesRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Authenticate token
	if request.Token != config.SettingsObj.AuthReadToken {
		http.Error(w, "Incorrect Token!", http.StatusUnauthorized)
		return
	}

	pastBatches := request.PastBatches
	if pastBatches < 0 {
		http.Error(w, "Past batches should be at least 0", http.StatusBadRequest)
		return
	}

	keys := redis.RedisClient.Keys(context.Background(), fmt.Sprintf("%s.commit_submission_batch.*", pkgs.ProcessTriggerKey)).Val()

	var logs []map[string]interface{}

	for _, key := range keys {
		entry, err := redis.Get(context.Background(), key)
		if err != nil {
			continue
		}

		var logEntry map[string]interface{}
		err = json.Unmarshal([]byte(entry), &logEntry)
		if err != nil {
			continue
		}

		logs = append(logs, logEntry)
	}

	if pastBatches > 0 && len(logs) > pastBatches {
		logs = logs[len(logs)-pastBatches:]
	}

	response := struct {
		Info struct {
			Success bool                     `json:"success"`
			Logs    []map[string]interface{} `json:"logs"`
		} `json:"info"`
		RequestID string `json:"request_id"`
	}{
		Info: struct {
			Success bool                     `json:"success"`
			Logs    []map[string]interface{} `json:"logs"`
		}{
			Success: true,
			Logs:    logs,
		},
		RequestID: r.Context().Value("request_id").(string),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func handleBatchResubmissions(w http.ResponseWriter, r *http.Request) {
	var request PastBatchesRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Authenticate token
	if request.Token != config.SettingsObj.AuthReadToken {
		http.Error(w, "Incorrect Token!", http.StatusUnauthorized)
		return
	}

	pastBatches := request.PastBatches
	if pastBatches < 0 {
		http.Error(w, "Past batches should be at least 0", http.StatusBadRequest)
		return
	}

	keys := redis.RedisClient.Keys(context.Background(), fmt.Sprintf("%s.batch_resubmission.*", pkgs.ProcessTriggerKey)).Val()

	var logs []map[string]interface{}

	for _, key := range keys {
		entry, err := redis.Get(context.Background(), key)
		if err != nil {
			continue
		}

		var logEntry map[string]interface{}
		err = json.Unmarshal([]byte(entry), &logEntry)
		if err != nil {
			continue
		}

		logs = append(logs, logEntry)
	}

	if pastBatches > 0 && len(logs) > pastBatches {
		logs = logs[len(logs)-pastBatches:]
	}

	response := struct {
		Info struct {
			Success bool                     `json:"success"`
			Logs    []map[string]interface{} `json:"logs"`
		} `json:"info"`
		RequestID string `json:"request_id"`
	}{
		Info: struct {
			Success bool                     `json:"success"`
			Logs    []map[string]interface{} `json:"logs"`
		}{
			Success: true,
			Logs:    logs,
		},
		RequestID: r.Context().Value("request_id").(string),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func handleIncludedEpochSubmissionsCount(w http.ResponseWriter, r *http.Request) {
	var request PastEpochsRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Authenticate token
	if request.Token != config.SettingsObj.AuthReadToken {
		http.Error(w, "Incorrect Token!", http.StatusUnauthorized)
		return
	}

	pastEpochs := request.PastEpochs
	if pastEpochs < 0 {
		http.Error(w, "Past epochs should be at least 0", http.StatusBadRequest)
		return
	}

	keys := redis.RedisClient.Keys(context.Background(), fmt.Sprintf("%s.build_batch.*", pkgs.ProcessTriggerKey)).Val()

	var totalSubmissions int

	for _, key := range keys {
		entry, err := redis.Get(context.Background(), key)
		if err != nil {
			continue
		}

		var logEntry map[string]interface{}
		err = json.Unmarshal([]byte(entry), &logEntry)
		if err != nil {
			continue
		}

		if epochIdStr, ok := logEntry["epoch_id"].(string); ok {
			epochId, err := strconv.Atoi(epochIdStr)
			if err != nil {
				continue
			}
			if pastEpochs == 0 || epochId >= pastEpochs {
				if count, ok := logEntry["submissions_count"].(float64); ok {
					totalSubmissions += int(count)
				}
			}
		}
	}

	response := struct {
		Info struct {
			Success          bool `json:"success"`
			TotalSubmissions int  `json:"total_submissions"`
		} `json:"info"`
		RequestID string `json:"request_id"`
	}{
		Info: struct {
			Success          bool `json:"success"`
			TotalSubmissions int  `json:"total_submissions"`
		}{
			Success:          true,
			TotalSubmissions: totalSubmissions,
		},
		RequestID: r.Context().Value("request_id").(string),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func handleReceivedEpochSubmissionsCount(w http.ResponseWriter, r *http.Request) {
	var request PastEpochsRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Authenticate token
	if request.Token != config.SettingsObj.AuthReadToken {
		http.Error(w, "Incorrect Token!", http.StatusUnauthorized)
		return
	}

	pastEpochs := request.PastEpochs
	if pastEpochs < 0 {
		http.Error(w, "Past epochs should be at least 0", http.StatusBadRequest)
		return
	}

	keys := redis.RedisClient.Keys(context.Background(), fmt.Sprintf("%s.*", pkgs.EpochSubmissionsCountKey)).Val()

	var totalSubmissions int

	for _, key := range keys {
		entry, err := redis.Get(context.Background(), key)
		if err != nil {
			continue
		}

		epochId, err := strconv.Atoi(strings.Split(key, ".")[1])
		if err != nil {
			continue
		}
		if pastEpochs == 0 || epochId >= pastEpochs {
			if count, err := strconv.Atoi(entry); err == nil {
				totalSubmissions += int(count)
			}
		}
	}

	response := struct {
		Info struct {
			Success          bool `json:"success"`
			TotalSubmissions int  `json:"total_submissions"`
		} `json:"info"`
		RequestID string `json:"request_id"`
	}{
		Info: struct {
			Success          bool `json:"success"`
			TotalSubmissions int  `json:"total_submissions"`
		}{
			Success:          true,
			TotalSubmissions: totalSubmissions,
		},
		RequestID: r.Context().Value("request_id").(string),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func handleReceivedEpochSubmissions(w http.ResponseWriter, r *http.Request) {
	var request PastEpochsRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Authenticate token
	if request.Token != config.SettingsObj.AuthReadToken {
		http.Error(w, "Incorrect Token!", http.StatusUnauthorized)
		return
	}

	pastEpochs := request.PastEpochs
	if pastEpochs < 0 {
		http.Error(w, "Past epochs should be at least 0", http.StatusBadRequest)
		return
	}

	// Fetch keys for epoch submissions
	keys := redis.RedisClient.Keys(context.Background(), fmt.Sprintf("%s.*", pkgs.EpochSubmissionsKey)).Val()

	var logs []map[string]interface{}

	for _, key := range keys {
		epochId := strings.TrimPrefix(key, fmt.Sprintf("%s.", pkgs.EpochSubmissionsKey))
		epochLog := make(map[string]interface{})
		epochLog["epoch_id"] = epochId

		submissions := redis.RedisClient.HGetAll(context.Background(), key).Val()
		submissionsMap := make(map[string]interface{})
		for submissionId, submissionData := range submissions {
			var submission interface{}
			if err := json.Unmarshal([]byte(submissionData), &submission); err != nil {
				continue
			}
			submissionsMap[submissionId] = submission
		}
		epochLog["submissions"] = submissionsMap

		logs = append(logs, epochLog)
	}

	if pastEpochs > 0 && len(logs) > pastEpochs {
		logs = logs[len(logs)-pastEpochs:]
	}

	response := struct {
		Info struct {
			Success bool                     `json:"success"`
			Logs    []map[string]interface{} `json:"logs"`
		} `json:"info"`
		RequestID string `json:"request_id"`
	}{
		Info: struct {
			Success bool                     `json:"success"`
			Logs    []map[string]interface{} `json:"logs"`
		}{
			Success: true,
			Logs:    logs,
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
	mux.HandleFunc("/triggeredCollectionFlows", handleTriggeredCollectionFlows)
	mux.HandleFunc("/finalizedBatchSubmissions", handleFinalizedBatchSubmissions)
	mux.HandleFunc("/builtBatches", handleBuiltBatches)
	mux.HandleFunc("/committedSubmissionBatches", handleCommittedSubmissionBatches)
	mux.HandleFunc("/batchResubmissions", handleBatchResubmissions)
	mux.HandleFunc("/receivedEpochSubmissions", handleReceivedEpochSubmissions)
	mux.HandleFunc("/receivedEpochSubmissionsCount", handleReceivedEpochSubmissionsCount)
	mux.HandleFunc("/includedEpochSubmissionsCount", handleIncludedEpochSubmissionsCount)
	// also add submission body per epoch
	handler := RequestMiddleware(mux)

	log.Println("Server is running on port 9988")
	log.Fatal(http.ListenAndServe(":9988", handler))
}
