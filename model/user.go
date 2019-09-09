package model

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	UserName string
	Password string
	NickName string
	Status string
	Avatar string `gorm:"size:1000"`
}

