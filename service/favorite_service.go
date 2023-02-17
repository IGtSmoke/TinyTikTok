package service

import (
	"TinyTikTok/model/dto"
)

// FavoriteService 赞接口
type FavoriteService interface {
	//Thumb 赞操作
	Thumb(int64, int64, int8) (dto.Result, error)
	//List 喜欢列表
	List(int64, string) (dto.FavoriteListResponse, error)
}
