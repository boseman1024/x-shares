package model

import "github.com/jinzhu/gorm"

type Comment struct {
	gorm.Model
	ShareRef    uint   `json:"shareref,omitempty"`
	CreatorRef  uint   `json:"creatorref,omitempty"`
	ReplyRef    uint   `json:"replyref,omitempty"`
	Creator     User   `gorm:"ForeignKey:CreatorRef;" json:"creator"`
	Description string `json:"description"`
}
