package service

import (
	"TinyTikTok/model/dto"
)

// CommentService 评论接口
type CommentService interface {
	//Comment 评论操作
	Comment(int64, int64, string) (dto.CommentActionResponse, error)
	//List 评论列表
	List(int64, int64) (dto.CommentListResponse, error)
}
