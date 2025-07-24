package config

import (
	"os"
	"strconv"

	log "github.com/sirupsen/logrus"
	logging "github.com/ipfs/go-log/v2"
)

var SettingsObj *Settings

type Settings struct {
	ClientUrl              string
	ContractAddress        string
	RedisHost              string
	RedisPort              string
	SlackReportingUrl      string
	DataMarketAddress      string
	RedisDB                string
	BootstrapPeers         string
	RendezvousPoint        string
	ListenerP2PPort        string
	PublicIP               string
	ABIFilePath     string
	OnlyEpoch0      bool
	AdvertiseRetries int
	AdvertiseRetryDelaySec int
}

func LoadConfig() {
	// Set libp2p logging level based on environment variable
	libp2pLogLevel := os.Getenv("LIBP2P_LOGGING")
	if libp2pLogLevel != "" {
		switch libp2pLogLevel {
		case "debug":
			logging.SetAllLoggers(logging.LevelDebug)
		case "info":
			logging.SetAllLoggers(logging.LevelInfo)
		case "warn":
			logging.SetAllLoggers(logging.LevelWarn)
		case "error":
			logging.SetAllLoggers(logging.LevelError)
		case "fatal":
			logging.SetAllLoggers(logging.LevelFatal)
		default:
			log.Warnf("Unknown LIBP2P_LOGGING level: %s. Defaulting to info.", libp2pLogLevel)
			logging.SetAllLoggers(logging.LevelInfo)
		}
	} else {
		// Default libp2p logging to info if not specified
		logging.SetAllLoggers(logging.LevelInfo)
	}

	config := Settings{
		ClientUrl:              getEnv("PROST_RPC_URL", ""),
		ContractAddress:        getEnv("PROTOCOL_STATE_CONTRACT", ""),
		RedisHost:              getEnv("REDIS_HOST", "localhost"),
		RedisPort:              getEnv("REDIS_PORT", "6379"),
		SlackReportingUrl:      getEnv("SLACK_REPORTING_URL", ""),
		DataMarketAddress:      getEnv("DATA_MARKET_ADDRESS", ""),
		RedisDB:                getEnv("REDIS_DB", "0"),
		BootstrapPeers:         getEnv("BOOTSTRAP_PEERS", ""),
		RendezvousPoint:        getEnv("RENDEZVOUS_POINT", ""),
		ListenerP2PPort:        getEnv("LISTENER_P2P_PORT", "4001"),
		PublicIP:               getEnv("PUBLIC_IP", ""),
		ABIFilePath:     getEnv("ABI_FILE_PATH", "abis/PowerloomProtocolState.json"),
		OnlyEpoch0:      getEnv("ONLY_EPOCH_0", "false") == "true",
		AdvertiseRetries:       getEnvAsInt("ADVERTISE_RETRIES", 5),
		AdvertiseRetryDelaySec: getEnvAsInt("ADVERTISE_RETRY_DELAY_SEC", 5),
	}

	// Check for any missing required environment variables and log errors
	missingEnvVars := []string{}
	if config.ClientUrl == "" {
		missingEnvVars = append(missingEnvVars, "PROST_RPC_URL")
	}
	if config.ContractAddress == "" {
		missingEnvVars = append(missingEnvVars, "PROTOCOL_STATE_CONTRACT")
	}
	
	if config.DataMarketAddress == "" {
		missingEnvVars = append(missingEnvVars, "DATA_MARKET_ADDRESS")
	}
	if config.RedisDB == "" {
		missingEnvVars = append(missingEnvVars, "REDIS_DB")
	}

	if len(missingEnvVars) > 0 {
		log.Fatalf("Missing required environment variables: %v", missingEnvVars)
	}

		checkOptionalEnvVar(config.SlackReportingUrl, "SLACK_REPORTING_URL")
	

	SettingsObj = &config
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func getEnvAsInt(key string, defaultValue int) int {
	valueStr := os.Getenv(key)
	if valueStr == "" {
		return defaultValue
	}
	value, err := strconv.Atoi(valueStr)
	if err != nil {
		log.Warnf("Invalid integer value for environment variable %s: %s. Using default value %d", key, valueStr, defaultValue)
		return defaultValue
	}
	return value
}

func checkOptionalEnvVar(value, key string) {
	if value == "" {
		log.Warnf("Optional environment variable %s is not set", key)
	}
}
