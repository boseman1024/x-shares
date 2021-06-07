package service

import (
	"github.com/jinzhu/gorm"
	"shares/db"
	"shares/model"
	"shares/serializer"
)

type FollowService struct {
	UserId      uint `json:"userId"`
	FollowingId uint `json:"followingId"`
	Status      string
}

func (service *FollowService) List() *serializer.Response {
	var group []model.Follow
	follow := model.Follow{
		UserId:      service.UserId,
		FollowingId: service.FollowingId,
		Status:      "1",
	}
	err := db.DB.Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Select("id,nickname,note,avatar").Unscoped()
	}).Preload("Follower", func(db *gorm.DB) *gorm.DB {
		return db.Select("id,nickname,note,avatar").Unscoped()
	}).Where(follow).Find(&group).Error
	if err != nil {
		return &serializer.Response{
			Code: 302,
			Msg:  "获取关注列表失败",
		}
	}
	return &serializer.Response{
		Code: 200,
		Data: group,
	}
}
func (service *FollowService) Get() (model.Follow, *serializer.Response) {
	follow := model.Follow{
		UserId:      service.UserId,
		FollowingId: service.FollowingId,
	}
	if err := db.DB.Where(follow).First(&follow).Error; err != nil {
		return follow, &serializer.Response{
			Code: 303,
			Msg:  "获取关注失败",
		}
	}
	return follow, nil
}
func (service *FollowService) Insert() (model.Follow, *serializer.Response) {
	follow := model.Follow{
		UserId:      service.UserId,
		FollowingId: service.FollowingId,
		Status:      "1",
	}
	if err := db.DB.Create(&follow).Error; err != nil {
		return follow, &serializer.Response{
			Code: 303,
			Msg:  "添加关注失败",
		}
	}
	return follow, nil
}
func (service *FollowService) Update(id uint, status string) (model.Follow, *serializer.Response) {
	follow := model.Follow{
		Model: gorm.Model{
			ID: id,
		},
		Status: status,
	}
	if err := db.DB.Save(&follow).Error; err != nil {
		return follow, &serializer.Response{
			Code: 304,
			Msg:  "修改关注失败",
		}
	}
	return follow, nil
}
