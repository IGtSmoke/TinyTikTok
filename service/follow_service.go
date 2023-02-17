// Package service contains interfaces for services
package service

import (
	"TinyTikTok/model/dto"
)

// FollowService 关注接口
type FollowService interface {
	// FollowUser 关注操作
	FollowUser(myId int64, userId int64) (dto.Result, error)
	// UnFollowUser 取关操作
	UnFollowUser(myId int64, userId int64) (dto.Result, error)
}
