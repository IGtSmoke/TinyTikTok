// Package dto contains DTO struct
package dto

import "time"

type CommentDTO struct {
	UserID      int64     // 评论用户id
	VideoID     int64     // 视频id
	CommentText string    // 评论内容
	CreateDate  time.Time // 评论发布的日期mm-dd
}
