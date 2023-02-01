package main

import (
	"fmt"
	"log"
	"url-healthcheck/config"
	"url-healthcheck/internal/grpc"
	"url-healthcheck/pkg/db"
	"url-healthcheck/pkg/db/mysqldb"
)

func main() {
	var err error

	config.LoadConfig()

	var db db.DB
	db, err = mysqldb.NewConnection(
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
	fmt.Println("database ready")

	fmt.Println("GRPCServer starting")
	log.Fatalf("failed to GRPCServer: %v", grpc.StartGRPCServer(db, config.GRPCPort))
}
