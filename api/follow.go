package api

import (
	"github.com/gin-gonic/gin"
	"shares/model"
	"shares/serializer"
	"shares/service"
)

func FollowList(c *gin.Context) {
	var followService service.FollowService
	if err := c.ShouldBindJSON(&followService); err == nil {
		res := followService.List()
		c.JSON(200, res)
	} else {
		c.JSON(200, serializer.Response{
			Code: 201,
			Msg:  "关注服务初始化失败",
		})
	}

}
func FollowChange(c *gin.Context) {
	var followService service.FollowService
	if err := c.ShouldBindJSON(&followService); err == nil {
		temp, _ := followService.Get()
		var follow model.Follow
		var err *serializer.Response
		if temp.ID != 0 {
			if temp.Status == "1" {
				temp.Status = "0"
			} else {
				temp.Status = "1"
			}
			follow, err = followService.Update(temp.ID, temp.Status)
		} else {
			follow, err = followService.Insert()
		}
		if err != nil {
			c.JSON(200, err)
		} else {
			c.JSON(200, serializer.Response{
				Code: 200,
				Data: follow,
			})
		}
	} else {
		c.JSON(200, serializer.Response{
			Code: 201,
			Msg:  "关注服务初始化失败",
		})
	}
}
