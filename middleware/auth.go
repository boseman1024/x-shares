package middleware

import (
	"github.com/gin-gonic/gin"
	"shares/serializer"
	"shares/util"
)

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if token == "" {
			c.JSON(200, serializer.Response{
				Code: 401,
				Msg:  "请求未携带token，无权限访问",
			})
			c.Abort()
			return
		}
		jwtUtil := util.JwtUtil{}
		if claims := jwtUtil.ParseToekn(token); claims == nil {
			c.JSON(200, serializer.Response{
				Code: 403,
				Msg:  "携带token无效",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}

func GetCurrentUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
	}
}
