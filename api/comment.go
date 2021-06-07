package api

import (
	"github.com/gin-gonic/gin"
	"shares/serializer"
	"shares/service"
)

func CommentList(c *gin.Context) {
	var commentService service.CommentService
	if err := c.ShouldBindJSON(&commentService); err == nil {
		res := commentService.List()
		c.JSON(200, res)
	} else {
		c.JSON(200, serializer.Response{
			Code: 201,
			Msg:  "评论服务初始化失败",
		})
	}
}

func CommentAdd(c *gin.Context) {
	var commentService service.CommentService
	if err := c.ShouldBindJSON(&commentService); err == nil {
		if comment, err := commentService.Insert(); err != nil {
			c.JSON(200, err)
		} else {
			c.JSON(200, serializer.Response{
				Code: 200,
				Data: comment,
			})
		}
	} else {
		c.JSON(200, serializer.Response{
			Code: 201,
			Msg:  "评论服务初始化失败",
		})
	}
}

func CommentDel(c *gin.Context) {
	var commentService service.CommentService
	if num, err := commentService.Delete(c.Param("id")); err != nil {
		c.JSON(200, err)
	} else {
		c.JSON(200, serializer.Response{
			Code: 200,
			Data: num,
		})
	}
}
