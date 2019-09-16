package api

import (
	"github.com/gin-gonic/gin"
	"shares/serializer"
	"shares/service"
)

func UserRegister(c *gin.Context){
	var userRegisterService service.UserRegisterService
	if err:=c.ShouldBindJSON(&userRegisterService);err==nil{
		if user,err := userRegisterService.Register();err!=nil{
			c.JSON(200,err)
		}else{
			c.JSON(200,serializer.Response{
				Code: 100,
				Data: user,
			})
		}
	}else{
		c.JSON(200,serializer.Response{
			Code: 101,
			Msg:    "注册服务初始化失败",
		})
	}
}

func UserLogin(c *gin.Context){
	var userLoginService service.UserLoginService
	if err:=c.ShouldBind(&userLoginService);err==nil{
		if token,user,err := userLoginService.Login();err!=nil{
			c.JSON(200,err)
		}else{
			c.JSON(200,serializer.Response{
				Code:110,
				Data:serializer.LoginResponse{
					Token:token,
					User:user,
				},
			})
		}
	}else{
		c.JSON(200,serializer.Response{
			Code: 111,
			Msg:    "登录服务初始化失败",
		})
	}
}

func UserLogout(c *gin.Context){

}

func UserMe(c *gin.Context){

}

func UserUpdate(c *gin.Context){

}

