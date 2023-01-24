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
	r := gin.Default()

	//以下的接口，都使用RefreshTokenInterceptor()中间件身份验证
	r.Use(utils.RefreshTokenInterceptor())
	routers.InitRouter(r)

	//依赖加载
	initDeps()
	err := r.Run()
	if err != nil {
		panic("无法启动项目:r.Run失败")
	}
}

// 依赖加载
func initDeps() {
	err := conf.LoadConfig()
	if err != nil {
		log.Err(err)
	}
	//初始化redis连接
	setup.Redis()
	setup.Gorm()
	setup.Minio()
}
