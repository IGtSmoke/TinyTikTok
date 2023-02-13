// Package main as the entry point of the program
package main

import (
	"TinyTikTok/conf"
	"TinyTikTok/conf/setup"
	"TinyTikTok/routers"
	"TinyTikTok/utils"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func main() {
	ginServer := gin.New()

	// 依赖加载
	initDeps()

	// 使用RefreshTokenInterceptor()中间件身份验证
	ginServer.Use(utils.SlowLog(), utils.RefreshTokenInterceptor(), gin.Recovery())
	routers.InitRouter(ginServer)
	if err := ginServer.Run(":" + conf.Conf.GinPort); err != nil {
		panic("无法启动项目:ginServer.Run失败")
	}
}

// 依赖加载
func initDeps() {
	if err := conf.LoadConfig(); err != nil {
		log.Err(err).Send()
	}
	setup.Zerolog("slow", "common")
	// 初始化redis连接
	setup.Redis()
	setup.Gorm()
	setup.Minio()
	setup.Snowflake()
}
