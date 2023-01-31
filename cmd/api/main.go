package main

import (
	"fmt"
	"log"
	"url-healthcheck/config"
	"url-healthcheck/internal/grpc"
	"url-healthcheck/pkg/db"
	"url-healthcheck/pkg/db/mysqldb"

	"github.com/joho/godotenv"
)

func main() {
	var err error

	err = config.LoadConfig()
	if err != nil {
		log.Fatalf("%v", err)
	}

	err = godotenv.Load()
	if err != nil {
		log.Fatalf("%v", err)
	}

	var conn db.DB
	conn, err = mysqldb.NewConnection(
		config.DBHost,
		config.DBPort,
		config.DBSchema,
		config.DBUser,
		config.DBPassword,
		config.DBDrop,
	)
	if err != nil {
		log.Fatalf("%v", err)
	}
	fmt.Printf("database ready")

	fmt.Println("GRPCServer starting")
	log.Fatalf("failed to GRPCServer: %v", grpc.StartGRPCServer(
		conn,
	))
}
