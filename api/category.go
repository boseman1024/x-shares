package api

import (
	"github.com/gin-gonic/gin"
	"shares/serializer"
	"shares/service"
)

//目录列表
func CategoryList(c *gin.Context) {
	var categoryService service.CategoryService
	if err := c.ShouldBindJSON(&categoryService); err == nil {
		if categoryList, err := categoryService.List(); err != nil {
			c.JSON(200, err)
		} else {
			c.JSON(200, serializer.Response{
				Code: 200,
				Data: categoryList,
			})
		}
	} else {
		c.JSON(200, serializer.Response{
			Code: 201,
			Msg:  "目录服务初始化失败",
		})
	}
}

//目录添加
func CategoryAdd(c *gin.Context) {
	var categoryService service.CategoryService
	if err := c.ShouldBindJSON(&categoryService); err == nil {
		if category, err := categoryService.Insert(); err != nil {
			c.JSON(200, err)
		} else {
			c.JSON(200, serializer.Response{
				Code: 200,
				Data: category,
			})
		}
	} else {
		c.JSON(200, serializer.Response{
			Code: 201,
			Msg:  "目录服务初始化失败",
		})
	}
}

//目录修改
func CategoryEdit(c *gin.Context) {
	var categoryService service.CategoryService
	if err := c.ShouldBindJSON(&categoryService); err == nil {
		if category, err := categoryService.Update(); err != nil {
			c.JSON(200, err)
		} else {
			c.JSON(200, serializer.Response{
				Code: 200,
				Data: category,
			})
		}
	} else {
		c.JSON(200, serializer.Response{
			Code: 201,
			Msg:  "目录服务初始化失败",
		})
	}
}

//目录删除
func CategoryDel(c *gin.Context) {
	var categoryService service.CategoryService
	if err := c.ShouldBindJSON(&categoryService); err == nil {
		if num, err := categoryService.Delete(); err != nil {
			c.JSON(200, err)
		} else {
			c.JSON(200, serializer.Response{
				Code: 200,
				Data: num,
			})
		}
	} else {
		c.JSON(200, serializer.Response{
			Code: 201,
			Msg:  "目录服务初始化失败",
		})
	}
}
