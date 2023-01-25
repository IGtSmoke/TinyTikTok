package service

import "github.com/gin-gonic/gin"

// RelationService 社交接口
type RelationService interface {
	// Action 关系操作
	Action(c *gin.Context)
	// Follow 用户关注列表
	Follow
	// Follower 用户粉丝列表
	Follower
	// Friend 用户好友列表
	Friend
}

// Follow 用户关注列表
type Follow interface {
	// List 用户关注列表
	List(c *gin.Context)
}

// Follower 用户粉丝列表
type Follower interface {
	// List 用户粉丝列表
	List(c *gin.Context)
}

// Friend 用户好友列表
type Friend interface {
	// List 用户好友列表
	List(c *gin.Context)
}
