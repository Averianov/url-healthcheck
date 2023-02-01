package mysqldb

import (
	"fmt"
	"strings"
	"time"
	"url-healthcheck/pkg/db"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type connection struct {
	*gorm.DB
}

// NewSession make new DB session
func NewConnection(host, port, name, user, password string, dropDB bool) (conn connection, err error) {

	builder := strings.Builder{}
	builder.WriteString(user)
	builder.WriteString(":")
	builder.WriteString(password)
	builder.WriteString("@tcp(")
	builder.WriteString(host)
	builder.WriteString(":")
	builder.WriteString(port)
	builder.WriteString(")/")
	builder.WriteString(name)
	builder.WriteString("?charset=utf8mb4&parseTime=True&loc=Local")

	// Waiting for MySql DB to be available
	var client *gorm.DB
	for {
		if client, err = gorm.Open(
			mysql.Open(builder.String()),
			&gorm.Config{},
		); err != nil {
			//fmt.Printf("Wait access to DB, %s\n", err.Error())
			time.Sleep(10 * time.Second)
			continue
		} else {
			break
		}
	}

	conn = connection{client}

	// When need start application from clear DB
	if dropDB {
		conn.Migrator().DropTable(&db.Check{})
		fmt.Println("db was dropped")
	}

	conn.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&db.Check{})
	fmt.Println("db was created if not exist")

	return
}
