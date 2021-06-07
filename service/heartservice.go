package service

import (
	"github.com/jinzhu/gorm"
	"shares/db"
	"shares/model"
	"shares/serializer"
)

type HeartService struct {
	UserId  uint `json:"userId"`
	ShareId uint `json:"shareId"`
}

func (service *HeartService) List() *serializer.Response {
	var group []model.Heart
	heart := model.Heart{
		UserId:  service.UserId,
		ShareId: service.ShareId,
		Status:  "1",
	}
	err := db.DB.Preload("Share").Preload("Share.User", func(db *gorm.DB) *gorm.DB {
		return db.Select("id,nickname,note,avatar").Unscoped()
	}).Where(heart).Find(&group).Error
	if err != nil {
		return &serializer.Response{
			Code: 302,
			Msg:  "获取收藏列表失败",
		}
	}
	return &serializer.Response{
		Code: 200,
		Data: group,
	}
}
func (service *HeartService) Get() (model.Heart, *serializer.Response) {
	heart := model.Heart{
		UserId:  service.UserId,
		ShareId: service.ShareId,
	}
	if err := db.DB.Where(heart).First(&heart).Error; err != nil {
		return heart, &serializer.Response{
			Code: 303,
			Msg:  "获取收藏失败",
		}
	}
	return heart, nil
}
func (service *HeartService) Insert() (model.Heart, *serializer.Response) {
	heart := model.Heart{
		UserId:  service.UserId,
		ShareId: service.ShareId,
		Status:  "1",
	}
	if err := db.DB.Create(&heart).Error; err != nil {
		return heart, &serializer.Response{
			Code: 303,
			Msg:  "添加收藏失败",
		}
	}
	return heart, nil
}
func (service *HeartService) Update(id uint, status string) (model.Heart, *serializer.Response) {
	heart := model.Heart{
		Model: gorm.Model{
			ID: id,
		},
		Status: status,
	}
	if err := db.DB.Save(&heart).Error; err != nil {
		return heart, &serializer.Response{
			Code: 304,
			Msg:  "修改收藏失败",
		}
	}
	return heart, nil
}
