// Package utils contains utility functions
package utils

import (
	"github.com/gin-gonic/gin"
)

// LoginInterceptor 用户是否有访问权限(是否登录)
func LoginInterceptor() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		value, exists := ctx.Get("userId")
		if !exists {
			Fail(ctx, "无权限访问")
			ctx.Abort()

			return
		}

		if value == "" {
			Fail(ctx, "无权限访问")
			ctx.Abort()

			return
		}

		ctx.Next()
	}
}
