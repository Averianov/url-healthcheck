package config

import (
	"log"
	"os"
	"strconv"
	"strings"
)

func mustEnvStr(key string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	} else {
		log.Panicf("Environment variable %v must be set.", key)
		return ""
	}
}

func optionalEnvStr(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func optionalEnvStrList(key string, fallback []string) []string {
	if value, ok := os.LookupEnv(key); ok {
		if value != "" {
			slc := strings.Split(value, ",")
			for i := range slc {
				slc[i] = strings.TrimSpace(slc[i])
			}
			return slc
		}
		return []string{}
	}
	return fallback
}

func optionalEnvBool(key string, fallback bool) bool {
	if strValue, ok := os.LookupEnv(key); ok {
		boolValue, err := strconv.ParseBool(strValue)
		if err != nil {
			log.Panicf("Failed to parse environment variable as bool %v with value %v.", key, strValue)
			return fallback
		}
		return boolValue
	}
	return fallback
}

func optionalEnvInt(key string, fallback int) int {
	if strValue, ok := os.LookupEnv(key); ok {
		intValue, err := strconv.Atoi(strValue)
		if err != nil {
			log.Panicf("Failed to parse environment variable as int %v with value %v.", key, strValue)
			return fallback
		}
		return intValue
	}
	return fallback
}

func optionalEnvInt64(key string, fallback int64) int64 {
	if value, ok := os.LookupEnv(key); ok {
		int64Value, err := strconv.ParseInt(value, 10, 64) // str to int64
		if err != nil {
			log.Panicf("Failed to parse environment variable as int64 %v with value %v.", key, value)
			return fallback
		}
		return int64Value
	}
	return fallback
}

func mustEnvInt64(key string) int64 {
	if value, ok := os.LookupEnv(key); ok {
		int64Value, err := strconv.ParseInt(value, 10, 64) // str to int64
		if err != nil {
			log.Panicf("Failed to parse environment variable as int64 %v with value %v.", key, value)
			return -1
		}
		return int64Value
	} else {
		log.Panicf("Environment variable %v must be set.", key)
		return -1
	}
}
