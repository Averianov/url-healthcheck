package main

import (
	"fmt"
	"log"
	"time"
	"url-healthcheck/config"
	"url-healthcheck/internal/dispatcher"
	"url-healthcheck/pkg/db"
	"url-healthcheck/pkg/db/mysqldb"
)

func main() {
	var err error

	config.LoadConfig()

	fmt.Printf("# Try connect to database %s:%s\n", config.DBHost, config.DBPort)

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

	fmt.Println("# URLDispatcher starting")
	d := dispatcher.NewURLDispatcher(db)
	log.Fatalf("failed to URLDispatcher: %v", d.StartURLDispatcher(time.Duration(config.HealthCheckDuration), config.URLConfig))
}
