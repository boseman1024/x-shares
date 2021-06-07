package service

import (
	"github.com/jinzhu/gorm"
	"log"
	"shares/db"
	"shares/model"
	"shares/serializer"
)

type CommentService struct {
	Id          uint   `json:"id"`
	ShareRef    uint   `json:"shareRef"`
	CreatorRef  uint   `json:"creatorRef"`
	ReplyRef    uint   `json:"repyRef"`
	Description string `json:"description"`
}

func (service *CommentService) List() *serializer.Response {
	var comments []model.Comment
	err := db.DB.Preload("Creator", func(db *gorm.DB) *gorm.DB {
		return db.Select("id,nickname,note,avatar").Unscoped()
	}).Where("share_ref=?", service.ShareRef).Find(&comments).Error
	if err != nil {
		log.Println(err)
		return &serializer.Response{
			Code: 302,
			Msg:  "获取评论列表失败",
		}
	}
	return &serializer.Response{
		Code: 200,
		Data: comments,
	}
}

func (service *CommentService) Insert() (model.Comment, *serializer.Response) {
	comment := model.Comment{
		Description: service.Description,
		ShareRef:    service.ShareRef,
		CreatorRef:  service.CreatorRef,
		ReplyRef:    service.ReplyRef,
	}
	if err := db.DB.Create(&comment).Error; err != nil {
		return comment, &serializer.Response{
			Code: 303,
			Msg:  "添加评论失败",
		}
	}
	return comment, nil
}

func (service *CommentService) Delete(commentId string) (int, *serializer.Response) {
	if err := db.DB.Where("id=?", commentId).Delete(&model.Comment{}).Error; err != nil {
		return 0, &serializer.Response{
			Code: 305,
			Msg:  "删除评论失败",
		}
	}
	return 1, nil
}
