package model

import "github.com/jinzhu/gorm"

type Share struct {
	gorm.Model
	Img           string     `json:"img"`
	Category      Category   `gorm:"ForeignKey:CategoryRefer;" json:"category"`
	CategoryRefer uint       `json:"categoryrefer,omitempty"`
	ShareTag      []ShareTag `gorm:"ForeignKey:ShareRefer" json:"sharetag"`
	User          User       `gorm:"ForeignKey:UserRefer;" json:"user"`
	UserRefer     uint       `json:"userrefer,omitempty"`
}
