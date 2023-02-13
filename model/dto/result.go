package dto

// Result 默认结构体.
type Result struct {
	// 必须大写，才能序列化
	StatusCode int64  `json:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg  string `json:"status_msg"`  // 返回状态描述
}

// UserLoginResponse 用户登录返回结构体.
type UserLoginResponse struct {
	Result
	UserID int64  `json:"user_id"`
	Token  string `json:"token"`
}

// UserInfoResponse 用户信息返回结构体.
type UserInfoResponse struct {
	Result
	User User `json:"user"` // 用户信息
}

// PublishListResponse 视频列表返回结构体.
type PublishListResponse struct {
	Result
	VideoList []Video `json:"video_list"` // 用户发布的视频列表
}

type PublishFeedResponse struct {
	NextTime *int64 `json:"next_time"` // 本次返回的视频中，发布最早的时间，作为下次请求时的latest_time
	Result
	VideoList []Video `json:"video_list"` // 视频列表
}

// User 组装响应体user信息
type User struct {
	FollowCount   int64  `json:"follow_count"`   // 关注总数
	FollowerCount int64  `json:"follower_count"` // 粉丝总数
	ID            int64  `json:"id"`             // 用户id
	IsFollow      bool   `json:"is_follow"`      // true-已关注，false-未关注
	Name          string `json:"name"`           // 用户名称
}

// Video 组装响应体需要的video信息
type Video struct {
	Author        User  `json:"author"`         // 视频作者信息
	CommentCount  int64 `json:"comment_count"`  // 视频的评论总数
	FavoriteCount int64 `json:"favorite_count"` // 视频的点赞总数
	IsFavorite    bool  `json:"is_favorite"`    // true-已点赞，false-未点赞
	VideoDTO
}
