package service

import "github.com/gin-gonic/gin"

// RelationService 社交接口
type RelationService interface {
	// Follow 用户关注列表
	FollowList(c *gin.Context)
	// Follower 用户粉丝列表
	FollowerList(c *gin.Context)
	// Friend 用户好友列表
	FriendList(c *gin.Context)
}
