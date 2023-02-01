package db

import (
	"gorm.io/gorm"
)

type DB interface {
	CreateCheck(chk *Check) (err error)
	GetCheckList() (chks []Check, err error)
}

type Check struct {
	gorm.Model
	ID      uint64 `gorm:"primaryKey;autoIncrement,unique;not null"`
	Url     string `gorm:"not null"`
	Type    string `gorm:"not null"`
	Status  string `gorm:"not null"`
	Comment string
}
