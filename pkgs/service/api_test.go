package service

import (
	"Listen/config"
	"Listen/pkgs/prost"
	"Listen/pkgs/redis"
	"Listen/pkgs/utils"
	"bytes"
	"context"
	"encoding/json"
	"github.com/alicebob/miniredis/v2"
	redisv8 "github.com/go-redis/redis/v8"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"math/big"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
)

var mr *miniredis.Miniredis

func TestMain(m *testing.M) {
	var err error
	mr, err = miniredis.Run()
	if err != nil {
		log.Fatalf("could not start miniredis: %v", err)
	}

	// Initialize the config settings
	config.SettingsObj = &config.Settings{
		ChainID:                 11165,
		ContractAddress:         "0x10c5E2ee14006B3860d4FdF6B173A30553ea6333",
		ClientUrl:               "https://rpc-prost-1h-sv2cc82c3x.t.conduit.xyz",
		SignerAccountAddressStr: "0x2132Bc41C5a9eAD0895B26cC2450FA3c35D07017",
		AuthReadToken:           "valid-token",
		RedisHost:               mr.Host(),
		RedisPort:               mr.Port(),
	}

	redis.RedisClient = redis.NewRedisClient()
	utils.InitLogger()

	prost.ConfigureClient()
	prost.ConfigureContractInstance()

	m.Run()

	mr.Close()
}

func setupRedisForGetTotalRewardsTest(t *testing.T, slotID int, totalRewards string) {
	t.Helper()

	// Create a mini Redis server for testing
	mr.FlushDB()

	redis.RedisClient = redisv8.NewClient(&redisv8.Options{
		Addr: mr.Addr(),
	})
	utils.InitLogger()

	// Populate Redis with test data
	key := redis.TotalSlotRewards(strconv.Itoa(slotID))
	err := redis.Set(context.Background(), key, totalRewards, 0)
	if err != nil {
		t.Fatalf("Failed to set test data in redis: %v", err)
	}
}

func TestHandleTotalRewards(t *testing.T) {
	tests := []struct {
		name           string
		slotID         int
		totalRewards   string
		expectedStatus int
		expectedReward int
	}{
		{
			name:           "Valid total rewards",
			slotID:         1,
			totalRewards:   "500",
			expectedStatus: http.StatusOK,
			expectedReward: 500,
		},
		{
			name:           "Empty total rewards",
			slotID:         2,
			totalRewards:   "",
			expectedStatus: http.StatusInternalServerError,
			expectedReward: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setupRedisForGetTotalRewardsTest(t, tt.slotID, tt.totalRewards)

			// Set up a valid request body
			requestBody := RewardsRequest{
				SlotID: tt.slotID,
				Token:  config.SettingsObj.AuthReadToken,
			}
			reqBody, err := json.Marshal(requestBody)
			if err != nil {
				t.Fatalf("could not marshal request body: %v", err)
			}

			// Create a new HTTP request
			req, err := http.NewRequest("POST", "/getTotalRewards", bytes.NewBuffer(reqBody))
			if err != nil {
				t.Fatalf("could not create HTTP request: %v", err)
			}

			// Create a ResponseRecorder to record the response
			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(handleTotalRewards)

			// Wrap the handler with the middleware
			testHandler := RequestMiddleware(handler)

			// Call the handler
			testHandler.ServeHTTP(rr, req)

			responseBody := rr.Body.String()
			t.Log("Response Body:", responseBody)

			// Check the status code
			if status := rr.Code; status != tt.expectedStatus {
				t.Errorf("handler returned wrong status code: got %v want %v", status, tt.expectedStatus)
			}

			// Parse the response body
			var response struct {
				Info struct {
					Success  bool `json:"success"`
					Response int  `json:"response"`
				} `json:"info"`
				RequestID string `json:"request_id"`
			}
			err = json.NewDecoder(rr.Body).Decode(&response)
			if err != nil {
				t.Fatalf("could not decode response body: %v", err)
			}

			// Validate the response
			if tt.expectedStatus == http.StatusOK && !response.Info.Success {
				t.Errorf("response success should be true")
			}
			if response.Info.Response != tt.expectedReward {
				t.Errorf("response rewards should match expected value: got %v want %v", response.Info.Response, tt.expectedReward)
			}
		})
	}
}

func setupRedisForDailyRewardsTest(t *testing.T, slotID int, day int, dailySubmissionCount string) {
	t.Helper()

	// Create a mini Redis server for testing
	mr.FlushDB()

	redis.RedisClient = redisv8.NewClient(&redisv8.Options{
		Addr: mr.Addr(),
	})
	utils.InitLogger()

	// Mock daily snapshot quota and reward base points
	dailySnapshotQuota = big.NewInt(10)
	rewardBasePoints = big.NewInt(1000)

	// Populate Redis with test data
	key := redis.SlotSubmissionKey(strconv.Itoa(slotID), strconv.Itoa(day))
	err := redis.Set(context.Background(), key, dailySubmissionCount, 0)
	if err != nil {
		t.Fatalf("Failed to set test data in redis: %v", err)
	}
}

func TestHandleDailyRewards(t *testing.T) {
	tests := []struct {
		name                  string
		slotID                int
		day                   int
		dailySubmissionCount  string
		expectedStatus        int
		expectedReward        int64
		expectedSuccessStatus bool
	}{
		{
			name:                  "Valid daily rewards",
			slotID:                1,
			day:                   10,
			dailySubmissionCount:  "15",
			expectedStatus:        http.StatusOK,
			expectedReward:        1000,
			expectedSuccessStatus: true,
		},
		{
			name:                  "Empty daily submission count",
			slotID:                2,
			day:                   10,
			dailySubmissionCount:  "",
			expectedStatus:        http.StatusInternalServerError,
			expectedReward:        0,
			expectedSuccessStatus: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setupRedisForDailyRewardsTest(t, tt.slotID, tt.day, tt.dailySubmissionCount)

			// Set up a valid request body
			requestBody := DailyRewardsRequest{
				SlotID: tt.slotID,
				Day:    tt.day,
				Token:  config.SettingsObj.AuthReadToken,
			}
			reqBody, err := json.Marshal(requestBody)
			if err != nil {
				t.Fatalf("could not marshal request body: %v", err)
			}

			// Create a new HTTP request
			req, err := http.NewRequest("POST", "/getDailyRewards", bytes.NewBuffer(reqBody))
			if err != nil {
				t.Fatalf("could not create HTTP request: %v", err)
			}

			// Create a ResponseRecorder to record the response
			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(handleDailyRewards)

			// Wrap the handler with the middleware
			testHandler := RequestMiddleware(handler)

			// Call the handler
			testHandler.ServeHTTP(rr, req)

			responseBody := rr.Body.String()
			t.Log("Response Body:", responseBody)

			// Check the status code
			if status := rr.Code; status != tt.expectedStatus {
				t.Errorf("handler returned wrong status code: got %v want %v", status, tt.expectedStatus)
			}

			// Parse the response body
			var response struct {
				Info struct {
					Success  bool `json:"success"`
					Response int  `json:"response"`
				} `json:"info"`
				RequestID string `json:"request_id"`
			}
			err = json.NewDecoder(rr.Body).Decode(&response)
			if err != nil {
				t.Fatalf("could not decode response body: %v", err)
			}

			// Validate the response
			if response.Info.Success != tt.expectedSuccessStatus {
				t.Errorf("response success status should match expected value: got %v want %v", response.Info.Success, tt.expectedSuccessStatus)
			}
			if int64(response.Info.Response) != tt.expectedReward {
				t.Errorf("response rewards should match expected value: got %v want %v", response.Info.Response, tt.expectedReward)
			}
		})
	}
}

func TestHandleTotalSubmissions(t *testing.T) {
	config.SettingsObj.AuthReadToken = "valid-token"
	//Day = big.NewInt(5)

	redis.Set(context.Background(), redis.SlotSubmissionKey("1", "5"), "100", 0)
	redis.Set(context.Background(), redis.SlotSubmissionKey("1", "4"), "120", 0)
	redis.Set(context.Background(), redis.SlotSubmissionKey("1", "3"), "80", 0)
	redis.Set(context.Background(), redis.SlotSubmissionKey("1", "2"), "400", 0)
	redis.Set(context.Background(), redis.SlotSubmissionKey("1", "1"), "25", 0)

	tests := []struct {
		name       string
		body       string
		statusCode int
		response   []DailySubmissions
	}{
		{
			name:       "Valid token, past days 1",
			body:       `{"slot_id": 1, "token": "valid-token", "past_days": 1}`,
			statusCode: http.StatusOK,
			response: []DailySubmissions{
				{Day: 5, Submissions: 100},
			},
		},
		{
			name:       "Valid token, past days 3",
			body:       `{"slot_id": 1, "token": "valid-token", "past_days": 3}`,
			statusCode: http.StatusOK,
			response: []DailySubmissions{
				{Day: 5, Submissions: 100},
				{Day: 4, Submissions: 120},
				{Day: 3, Submissions: 80},
			},
		},
		{
			name:       "Valid token, total submissions till date",
			body:       `{"slot_id": 1, "token": "valid-token", "past_days": 5}`,
			statusCode: http.StatusOK,
			response: []DailySubmissions{
				{Day: 5, Submissions: 100},
				{Day: 4, Submissions: 120},
				{Day: 3, Submissions: 80},
				{Day: 2, Submissions: 400},
				{Day: 1, Submissions: 25},
			},
		},
		{
			name:       "Valid token, negative past days",
			body:       `{"slot_id": 1, "token": "valid-token", "past_days": -1}`,
			statusCode: http.StatusBadRequest,
			response:   nil,
		},
		{
			name:       "Invalid token",
			body:       `{"slot_id": 1, "token": "invalid-token", "past_days": 1}`,
			statusCode: http.StatusUnauthorized,
			response:   nil,
		},
		{
			name:       "Invalid slot ID",
			body:       `{"slot_id": 10001, "token": "valid-token", "past_days": 1}`,
			statusCode: http.StatusBadRequest,
			response:   nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest("POST", "/totalSubmissions", strings.NewReader(tt.body))
			if err != nil {
				t.Fatal(err)
			}
			req.Header.Set("Content-Type", "application/json")

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(handleTotalSubmissions)
			testHandler := RequestMiddleware(handler)
			testHandler.ServeHTTP(rr, req)

			responseBody := rr.Body.String()
			t.Log("Response Body:", responseBody)

			assert.Equal(t, tt.statusCode, rr.Code)

			if tt.statusCode == http.StatusOK {
				var response struct {
					Info struct {
						Success  bool               `json:"success"`
						Response []DailySubmissions `json:"response"`
					} `json:"info"`
					RequestID string `json:"request_id"`
				}
				err = json.NewDecoder(rr.Body).Decode(&response)

				err := json.Unmarshal([]byte(responseBody), &response)
				assert.NoError(t, err)
				assert.Equal(t, tt.response, response.Info.Response)
			}
		})
	}
}
