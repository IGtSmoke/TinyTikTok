package service

import "github.com/gin-gonic/gin"

// PublishService 视频接口
type PublishService interface {
	// Action 视频投稿
	Action(c *gin.Context)
	// List 发布列表
	List(c *gin.Context)
}
