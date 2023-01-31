package main

import (
	"log"
	"url-healthcheck/config"
	"url-healthcheck/internal/dispatcher"
	"url-healthcheck/pkg/db"
	"url-healthcheck/pkg/db/mysqldb"
)

func main() {
	var err error

	err = config.LoadConfig()
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

	d := dispatcher.NewURLDispatcher(conn, "config", "url.json")
	log.Fatalf("failed to URLDispatcher: %v", d.StartURLDispatcher())
}
