// Package service contains interfaces for services
package service

import (
	"TinyTikTok/model/dto"

	"github.com/gin-gonic/gin"
)

// FollowService 关注接口
type FollowService interface {
	// List 用户关注列表
	List(c *gin.Context)
	// Follow 关注操作
	FollowUser(myId int64, userId int64) (dto.Result, error)
	// UnFollow 取关操作
	UnFollowUser(myId int64, userId int64) (dto.Result, error)
}
