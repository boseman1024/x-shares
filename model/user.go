package model

import (
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	gorm.Model
	Username string
	Password string
	Nickname string
	Status string
	Avatar string `gorm:"size:1000"`
}

const(
	//加密难度
	PasswordCost = 12

)

func (user *User) SetPwd(pwd string) error{
	bytes,err := bcrypt.GenerateFromPassword([]byte(pwd),PasswordCost)
	if err!=nil{
		return err
	} 
	user.Password = string(bytes)
	return nil
}

func (user *User) CheckPwd(pwd string) bool{
	err := bcrypt.CompareHashAndPassword([]byte(user.Password),[]byte(pwd))
	return err==nil
}