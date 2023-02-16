// Package dto contains DTO struct
package dto

// CommentDTO 评论信息
type CommentDTO struct {
	Id          uint   `gorm:"primarykey" json:"id"` // videoId由数据库生成
	UserID      int64  `json:"user_id"`              // 评论用户id
	VideoID     int64  `json:"video_id"`             // 视频id
	CommentText string `json:"content"`              // 评论内容
	CreateDate  string `json:"create_date"`          // 评论发布的日期mm-dd
}
