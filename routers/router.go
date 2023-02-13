// Package routers contains router functions
// Package routers contains router functions
package routers

import (
	"TinyTikTok/controller"
	"TinyTikTok/utils"
	"errors"

	"github.com/gin-gonic/gin"
)

// InitRouter 初始化路由
func InitRouter(ginServer *gin.Engine) {
	apiRouter := ginServer.Group("/douyin")
	// 基础接口
	apiRouter.GET("/feed/", controller.Feed)
	apiRouter.POST("/user/register/", controller.Register)
	apiRouter.POST("/user/login/", controller.Login)

	//以下api均需要鉴权
	ginServer.Use(utils.LoginInterceptor())
	apiRouter.POST("/publish/action/", controller.Action)
	apiRouter.GET("/publish/list/", controller.List)
	apiRouter.GET("/user/", controller.UserInfo)
	// 互动接口
	apiRouter.POST("/favorite/action/", controller.Thumb)
	apiRouter.GET("/favorite/list/")
	apiRouter.POST("/comment/action/")
	apiRouter.GET("/comment/list/")
	// 社交接口
	apiRouter.POST("/relation/action/", controller.Follow)
	apiRouter.GET("/relation/follow/list/", controller.FollowList)
	apiRouter.GET("/relation/follower/list", controller.FollowerList)
	apiRouter.GET("/relation/friend/list", controller.FriendList)

	// 404
	ginServer.NoRoute(func(c *gin.Context) {
		utils.Fail(c, errors.New("bad router"))
	})
}
