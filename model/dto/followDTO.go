package dto

// FollowDTO 关注信息
type FollowDTO struct {
	ID         int64
	UserID     int64
	FollowerID int64
	Cancel     int8
}
