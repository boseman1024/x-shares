package service

import (
	"encoding/json"
	"github.com/jinzhu/gorm"
	"log"
	"shares/db"
	"shares/model"
	"shares/serializer"
)

type ShareService struct {
	Id  uint
	Img string
}

type ShareListService struct {
	Uid uint `json:"uid"`
}

func (service *ShareListService) Index() *serializer.Response {
	var shares []model.Share
	err := db.DB.Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Select("id,nickname,note,avatar").Unscoped()
	}).Joins("left join follows on follows.following_id = user_refer").Where("follows.user_id=?", service.Uid).Order("created_at desc").Find(&shares).Error
	if err != nil {
		log.Println(err)
		return &serializer.Response{
			Code: 302,
			Msg:  "获取分享列表失败",
		}
	}
	return &serializer.Response{
		Code: 200,
		Data: shares,
	}
}
func (service *ShareListService) List() *serializer.Response {
	var shares []model.Share
	share := model.Share{
		UserRefer: service.Uid,
	}
	err := db.DB.Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Select("id,nickname,note,avatar").Unscoped()
	}).Preload("ShareTag").Where(share).Find(&shares).Error
	if err != nil {
		log.Println(err)
		return &serializer.Response{
			Code: 302,
			Msg:  "获取分享列表失败",
		}
	}
	return &serializer.Response{
		Code: 200,
		Data: shares,
	}
}
func (service *ShareService) Info() (model.Share, *serializer.Response) {
	share := model.Share{
		Model: gorm.Model{
			ID: service.Id,
		},
	}
	if err := db.DB.First(&share).Error; err != nil {
		return share, &serializer.Response{
			Code: 303,
			Msg:  "获取分享失败",
		}
	}
	return share, nil
}
func (service *ShareService) Insert(userId uint, tags string) (model.Share, *serializer.Response) {
	var sharetag []model.ShareTag
	json.Unmarshal([]byte(tags), &sharetag)
	share := model.Share{
		Img:       service.Img,
		UserRefer: userId,
		ShareTag:  sharetag,
	}
	if err := db.DB.Create(&share).Error; err != nil {
		return share, &serializer.Response{
			Code: 303,
			Msg:  "添加分享失败",
		}
	}
	return share, nil
}
func (service *ShareService) Update() (model.Share, *serializer.Response) {
	share := model.Share{
		Model: gorm.Model{
			ID: service.Id,
		},
		Img: service.Img,
	}
	if err := db.DB.Save(&share).Error; err != nil {
		return share, &serializer.Response{
			Code: 304,
			Msg:  "修改分享失败",
		}
	}
	return share, nil
}
func (service *ShareService) Delete(shareId string) (int, *serializer.Response) {
	if err := db.DB.Where("id=?", shareId).Delete(&model.Share{}).Error; err != nil {
		return 0, &serializer.Response{
			Code: 305,
			Msg:  "删除分享失败",
		}
	}
	return 1, nil
}
