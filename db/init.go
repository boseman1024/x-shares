package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"shares/model"
	"time"
)

var DB *gorm.DB

func Init(){
	db, err := gorm.Open("mysql", "root:123456@/shares?charset=utf8&parseTime=True&loc=Local")
	if err!=nil{
		fmt.Println("连接数据库失败")
		panic(err)
	}
	//日志
	db.LogMode(true)
	//连接池设置
	//最大空闲连接数
	db.DB().SetMaxIdleConns(20)
	//最大打开连接数
	db.DB().SetMaxOpenConns(100)
	//超时
	db.DB().SetConnMaxLifetime(time.Second*30)
	//建表
	db.AutoMigrate(&model.User{})

	DB = db
}