package router

import (
	"github.com/gin-gonic/gin"
	"shares/api"
	"shares/middleware"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, "pong")
	})

	v1 := r.Group("/api/v1")
	{
		//用户注册
		v1.POST("user/register", api.UserRegister)
		//用户登录
		v1.POST("user/login", api.UserLogin)
		//权限访问
		auth := v1.Group("/")
		auth.Use(middleware.AuthRequired())
		{
			auth.POST("index", api.ShareIndex)
			//用户
			auth.GET("user/get/:id", api.UserGet)
			auth.DELETE("user/logout", api.UserLogout)
			auth.POST("user/update", api.UserUpdate)
			//分类
			auth.POST("category/list", api.CategoryList)
			auth.POST("category/add", api.CategoryAdd)
			auth.POST("category/edit", api.CategoryEdit)
			auth.DELETE("category/del", api.CategoryDel)
			//分享
			auth.POST("share/list", api.ShareList)
			auth.GET("share/info", api.ShareInfo)
			auth.POST("share/add", api.ShareAdd)
			auth.GET("share/edit", api.ShareEdit)
			auth.DELETE("share/del/:id", api.ShareDel)
			//关注
			auth.POST("follow/list", api.FollowList)
			auth.POST("follow/change", api.FollowChange)
			//爱心
			auth.POST("heart/list", api.HeartList)
			auth.POST("heart/change", api.HeartChange)
			//评论
			auth.POST("comment/list", api.CommentList)
			auth.POST("comment/add", api.CommentAdd)
			auth.DELETE("comment/del/:id", api.CommentDel)
		}
	}
	return r
}
