package db

import (
	"gorm.io/gorm"
)

type DB interface {
	CreateCheck(chk *Check) (err error)
	GetCheckList() (chks []*Check, err error)
}

type Check struct {
	gorm.Model
	Id      string `json:"id" gorm:"unique;not null"`
	Url     string `json:"url" gorm:"unique;not null"`
	Type    string `json:"type" gorm:"not null"`
	Comment string `json:"comment" gorm:"not null"`
}
