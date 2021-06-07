package model

import "github.com/jinzhu/gorm"

type Heart struct {
	gorm.Model
	UserId  uint   `json:"userid"`
	ShareId uint   `json:"shareid"`
	Status  string `json:"status"`
	Share   Share  `gorm:"ForeignKey:ShareId;" json:"share"`
}
