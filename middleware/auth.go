package middleware

import (
	"github.com/gin-gonic/gin"
)

func AuthRequired() gin.HandlerFunc{
	return func(c *gin.Context){
		c.Next()
	}
}

func GetCurrentUser() gin.HandlerFunc{
	return func(c *gin.Context){
		c.Next()
	}
}