package service

import (
	"github.com/jinzhu/gorm"
	"shares/db"
	"shares/model"
	"shares/serializer"
)

type CategoryService struct {
	Id          uint
	Name        string
	Description string
	Pid         uint
}

func (service *CategoryService) List() ([]model.Category, *serializer.Response) {
	var categories []model.Category
	if err := db.DB.Find(&categories).Error; err != nil {
		return nil, &serializer.Response{
			Code: 202,
			Msg:  "获取目录列表失败",
		}
	}
	return categories, nil
}

func (service *CategoryService) Insert() (model.Category, *serializer.Response) {
	category := model.Category{
		Name:        service.Name,
		Description: service.Description,
		Pid:         service.Pid,
	}
	if err := db.DB.Create(&category).Error; err != nil {
		return category, &serializer.Response{
			Code: 203,
			Msg:  "添加目录失败",
		}
	}
	return category, nil
}

func (service *CategoryService) Update() (model.Category, *serializer.Response) {
	category := model.Category{
		Model: gorm.Model{
			ID: service.Id,
		},
		Name:        service.Name,
		Description: service.Description,
		Pid:         service.Pid,
	}
	if err := db.DB.Save(&category).Error; err != nil {
		return category, &serializer.Response{
			Code: 204,
			Msg:  "修改目录失败",
		}
	}
	return category, nil
}

func (service *CategoryService) Delete() (int, *serializer.Response) {
	if err := db.DB.Where("id=?", service.Id).Delete(&model.Category{}).Error; err != nil {
		return 0, &serializer.Response{
			Code: 205,
			Msg:  "删除目录失败",
		}
	}
	return 1, nil
}
