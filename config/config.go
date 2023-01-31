package config

import (
	"github.com/joho/godotenv"
)

var (
	DBHost              string
	DBPort              string
	DBSchema            string
	DBUser              string
	DBPassword          string
	DBDrop              bool
	HealthCheckDuration int64
)

const (
	DefaulDBtHost              = "localhost"
	DefaultDBPort              = "3306"
	DefaultHealthCheckDuration = 60 * 10 // in seconds (default 10 minutes)
)

func LoadConfig() (err error) {
	err = godotenv.Load(".env")
	if err != nil {
		err = nil
	}

	DBHost = optionalEnvStr("DB_HOSTS", DefaulDBtHost)
	DBPort = optionalEnvStr("DB_PORT", DefaultDBPort)
	DBSchema = mustEnvStr("DB_SCHEMA")
	DBUser = optionalEnvStr("DB_USER", "")
	DBPassword = optionalEnvStr("DB_PASSWORD", "")
	DBDrop = optionalEnvBool("DB_DROP", false)
	HealthCheckDuration = optionalEnvInt64("HCK_DURATION", DefaultHealthCheckDuration)

	return
}
