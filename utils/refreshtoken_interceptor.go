package utils

import (
	"TinyTikTok/conf/setup"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type Login struct {
	Token string `form:"token"  json:"token" uri:"token" xml:"token"`
}

// RefreshTokenInterceptor 刷新token(有token刷新,无token直接放过)
func RefreshTokenInterceptor() gin.HandlerFunc {
	return func(c *gin.Context) {

		var login Login
		err := c.ShouldBind(&login)
		if err != nil {
			log.Err(err)
		}
		//不存在token
		if login.Token == "" {
			c.Next()

			return
		}
		tokenKey := LoginUserKey + login.Token
		//取出userId
		userId, err := setup.Rdb.HGet(setup.Rctx, tokenKey, "userId").Result()
		if err != nil {
			log.Err(err)
			c.Next()

			return
		}
		if userId == "" {
			c.Next()

			return
		}
		//将userId存入context
		c.Set("userId", userId)
		//刷新token有效期
		setup.Rdb.Expire(setup.Rctx, tokenKey, LoginUserTTL)
		c.Next()
	}
}
