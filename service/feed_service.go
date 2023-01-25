package service

import "github.com/gin-gonic/gin"

// FeedService 视频流接口
type FeedService interface {
	// Feed 视频流
	Feed(c *gin.Context)
}
