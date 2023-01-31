package mysqldb

import (
	"fmt"
	"url-healthcheck/pkg/db"
)

// GetCheckByID function set Check Url to DB
func (db connection) CreateCheck(chk *db.Check) (err error) {
	db.Create(chk)
	if chk.ID <= 0 {
		err = fmt.Errorf("Failed to create check, connection error.")
		fmt.Printf("%v", err)
	}
	return
}

// GetCheckByID function get fromm DB Url checks
func (db connection) GetCheckList() (chks []*db.Check, err error) {
	err = db.Find(chks).Error
	if err != nil {
		fmt.Printf("%v", err)
	}
	return
}
