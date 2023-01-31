package dispatcher

import (
	"log"
	"url-healthcheck/internal/utils"
	"url-healthcheck/pkg/db"
)

type urlDispatcher struct {
	conn db.DB
	Urls []utils.URLs
}

func NewURLDispatcher(conn db.DB, configPath, configFile string) (c *urlDispatcher) {
	c = &urlDispatcher{
		conn: conn,
	}

	urls, err := utils.ReadJSON(configPath, configFile)
	if err != nil {
		log.Fatalf("godotenv failed to load json file: %v", err)
	}

	c.Urls = urls
	return
}

func (c *urlDispatcher) StartURLDispatcher() (err error) {
	//
	return
}
