package service

import (
	"TinyTikTok/model/dto"
	"mime/multipart"
)

// PublishService 视频接口
type PublishService interface {
	// Action 视频投稿
	Action(int64, string, *multipart.FileHeader) (dto.Result, error)
	// List 发布列表
	List(int64, string) (dto.PublishListResponse, error)
}
