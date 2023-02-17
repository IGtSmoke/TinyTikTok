package service

import (
	"TinyTikTok/model/dto"
)

// RelationService 社交接口
type RelationService interface {
	// ShowFollowList 用户关注列表
	ShowFollowList(int64, int64) (dto.RelationList, error)
	// ShowFollowerList 用户粉丝列表
	ShowFollowerList(int64, int64) (dto.RelationList, error)
	// ShowFriendList 用户好友列表
	ShowFriendList(int64, int64) (dto.RelationList, error)
}
