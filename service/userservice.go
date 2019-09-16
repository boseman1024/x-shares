package service

import (
	"shares/db"
	"shares/model"
	"shares/serializer"
	"shares/util"
)

type UserLoginService struct {
	Username string `from:"username" json:"username" binding:"required,min=5,max=30"`
	Password string `from:"password" json:"password" binding:"required,min=6,max=30"`
}
type UserRegisterService struct {
	Username        string `from:"username" json:"username" binding:"required,min=5,max=30"`
	Nickname        string `from:"nickname" json:"nickname" binding:"required,min=2,max=30"`
	Password        string `from:"password" json:"password" binding:"required,min=6,max=30"`
	PasswordConfirm string `from:"password_confirm" json:"password_confirm" binding:"required,min=6,max=30"`
}

func (service *UserLoginService) Login() (string, model.User, *serializer.Response) {
	var user model.User
	if err := db.DB.Where("username=?", service.Username).First(&user).Error; err != nil {
		return "", user, &serializer.Response{
			Code: 102,
			Msg:  "账号或密码错误",
		}
	}
	if isPwd := user.CheckPwd(service.Password); !isPwd {
		return "", user, &serializer.Response{
			Code: 103,
			Msg:  "账号或密码错误",
		}
	}
	jwtUtil := util.JwtUtil{}
	token, err := jwtUtil.CreateToken(user.ID, user.Username)
	if err != nil {
		return "", user, &serializer.Response{
			Code: 104,
			Msg:  "Token生成失败",
		}
	}
	user.Password = ""
	return token, user, nil
}

func (service *UserRegisterService) Register() (model.User, *serializer.Response) {
	user := model.User{
		Username: service.Username,
		Nickname: service.Nickname,
	}
	if err := service.Valid(); err != nil {
		return user, err
	}
	if err := user.SetPwd(service.Password); err != nil {
		return user, &serializer.Response{
			Code: 112,
			Msg:  "密码加密失败",
		}
	}
	if err := db.DB.Create(&user).Error; err != nil {
		return user, &serializer.Response{
			Code: 113,
			Msg:  "注册失败",
		}
	}
	return user, nil
}

func (service *UserRegisterService) Valid() *serializer.Response {
	if service.Password != service.PasswordConfirm {
		return &serializer.Response{
			Code: 114,
			Msg:  "两次输入的密码不匹配",
		}
	}
	var user model.User
	db.DB.Where("nickname=?", service.Nickname).First(&user)
	if user.Nickname != "" {
		return &serializer.Response{
			Code: 115,
			Msg:  "当前昵称已被占用",
		}
	}
	db.DB.Where("username=?", service.Username).First(&user)
	if user.Username != "" {
		return &serializer.Response{
			Code: 116,
			Msg:  "当前用户名已被占用",
		}
	}
	return nil
}
