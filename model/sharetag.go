package model

import "github.com/jinzhu/gorm"

type ShareTag struct {
	gorm.Model
	Name        string `json:"name"`
	Description string `json:"description"`
	Url         string `json:"url"`
	X           float32
	Y           float32
	ShareRefer  uint
}
