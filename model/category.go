package model

import "github.com/jinzhu/gorm"

type Category struct {
	gorm.Model
	Name        string
	Description string
	Pid         uint
}
