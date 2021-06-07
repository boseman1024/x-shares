package api

import (
	"github.com/gin-gonic/gin"
	"shares/model"
	"shares/serializer"
	"shares/service"
)

func HeartList(c *gin.Context) {
	var heartService service.HeartService
	if err := c.ShouldBindJSON(&heartService); err == nil {
		res := heartService.List()
		c.JSON(200, res)
	} else {
		c.JSON(200, serializer.Response{
			Code: 201,
			Msg:  "收藏服务初始化失败",
		})
	}

}
func HeartChange(c *gin.Context) {
	var heartService service.HeartService
	if err := c.ShouldBindJSON(&heartService); err == nil {
		temp, _ := heartService.Get()
		var heart model.Heart
		var err *serializer.Response
		if temp.ID != 0 {
			if temp.Status == "1" {
				temp.Status = "0"
			} else {
				temp.Status = "1"
			}
			heart, err = heartService.Update(temp.ID, temp.Status)
		} else {
			heart, err = heartService.Insert()
		}
		if err != nil {
			c.JSON(200, err)
		} else {
			c.JSON(200, serializer.Response{
				Code: 200,
				Data: heart,
			})
		}
	} else {
		c.JSON(200, serializer.Response{
			Code: 201,
			Msg:  "收藏服务初始化失败",
		})
	}
}
