package dto

import "time"

type CommentDTO struct {
	UserId      int64     //评论用户id
	VideoId     int64     //视频id
	CommentText string    //评论内容
	CreateDate  time.Time //评论发布的日期mm-dd
}
