// Package utils contains utility functions
package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

// LoginInterceptor 用户是否有访问权限(是否登录)
func LoginInterceptor() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userId, err := GetUserIdByMiddleware(ctx)
		if err != nil {
			log.Err(err).Msg("LoginInterceptor中userId不存在")
			ctx.Abort()
			return
		}
		if userId == 0 {
			Fail(ctx, "LoginInterceptor中userId不存在")
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}
