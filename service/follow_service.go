// Package service contains interfaces for services
package service

import "github.com/gin-gonic/gin"

// FollowService 关注接口
type FollowService interface {
	// List 用户关注列表
	List(c *gin.Context)
}
