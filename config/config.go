package config

import (
	"log"

	"github.com/joho/godotenv"
)

var (
	DBHost              string
	DBPort              string
	DBSchema            string
	DBUser              string
	DBPassword          string
	DBDrop              bool
	URLConfig           string
	GRPCPort            string
	HealthCheckDuration int64
)

const (
	DefaultDBPort              = "3306"
	DefaultDBSchema            = "urlcheck"
	DefaultDBUser              = "checker"
	DefaultDBPassword          = "checker"
	DefaultURLConfig           = "url.json"
	DefaultGRPCPort            = "443"
	DefaultHealthCheckDuration = 1 // in seconds (default 10 minutes)
)

// LoadConfig load enveronment form .env file if it exist and read enveronment
func LoadConfig() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println(".env not found")
	}

	DBHost = mustEnvStr("DB_HOST")
	DBPort = optionalEnvStr("DB_PORT", DefaultDBPort)
	DBSchema = optionalEnvStr("DB_SCHEMA", DefaultDBSchema)
	DBUser = optionalEnvStr("DB_USER", DefaultDBUser)
	DBPassword = optionalEnvStr("DB_PASSWORD", DefaultDBPassword)
	DBDrop = optionalEnvBool("DB_DROP", false)
	URLConfig = optionalEnvStr("URL_CONFIG", DefaultURLConfig)
	GRPCPort = optionalEnvStr("GRPC_PORT", DefaultGRPCPort)
	HealthCheckDuration = optionalEnvInt64("HCK_DURATION", DefaultHealthCheckDuration)

	return
}
