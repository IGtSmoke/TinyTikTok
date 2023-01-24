package routers

import (
	"TinyTikTok/utils"
	"github.com/gin-gonic/gin"
)

// InitRouter 初始化路由
func InitRouter(r *gin.Engine) {
	apiRouter := r.Group("/douyin")
	// 基础接口
	apiRouter.GET("/feed/")
	apiRouter.POST("/publish/action/")
	apiRouter.GET("/publish/list/")
	apiRouter.GET("/user/")
	apiRouter.POST("/user/register/")
	apiRouter.POST("/user/login/")
	// 互动接口
	apiRouter.POST("/favorite/action/")
	apiRouter.GET("/favorite/list/")
	apiRouter.POST("/comment/action/")
	apiRouter.GET("/comment/list/")
	// 社交接口
	apiRouter.POST("/relation/action/")
	apiRouter.GET("/relation/follow/list/")
	apiRouter.GET("/relation/follower/list")

	// 404
	r.NoRoute(func(c *gin.Context) {
		utils.Fail(c, "bad router")
	})
}
