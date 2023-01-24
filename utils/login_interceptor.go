package utils

import (
	"github.com/gin-gonic/gin"
)

// LoginInterceptor 用户是否有访问权限(是否登录)
func LoginInterceptor() gin.HandlerFunc {
	return func(c *gin.Context) {
		value, exists := c.Get("userId")
		if !exists {
			Fail(c, "无权限访问")
			c.Abort()
			return
		}

		if value == "" {
			Fail(c, "无权限访问")
			c.Abort()
			return
		}
		c.Next()
	}
}
