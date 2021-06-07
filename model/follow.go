package model

import "github.com/jinzhu/gorm"

type Follow struct {
	gorm.Model
	UserId      uint   `json:"userid"`
	FollowingId uint   `json:"followingid"`
	Status      string `json:"status"`
	User        User   `gorm:"ForeignKey:FollowingId;" json:"user"`
	Follower    User   `gorm:"ForeignKey:UserId;" json:"follower"`
}
