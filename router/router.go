package router

import (
	"github.com/gin-gonic/gin"
	"shares/api"
	"shares/middleware"
)

func NewRouter() *gin.Engine{
	r:=gin.Default()
	r.GET("/ping",func(c *gin.Context){
		c.JSON(200,"pong")
	})

	v1:=r.Group("/api/v1")
	{
		//用户注册
		v1.POST("user/register",api.UserRegister)
		//用户登录
		v1.POST("user/login",api.UserLogin)
		//权限访问
		auth:=v1.Group("/")
		auth.Use(middleware.AuthRequired())
		{
			auth.GET("user/me",api.UserMe)
			auth.DELETE("user/logout",api.UserLogout)
			auth.POST("user/update",api.UserUpdate)
		}
	}
	return r
}