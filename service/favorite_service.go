package service

import (
	"TinyTikTok/model/dto"
	"github.com/gin-gonic/gin"
)

// FavoriteService 赞接口
type FavoriteService interface {
	//Thumb 赞操作
	Thumb(int64, int64, bool) (dto.Result, error)
	//List 喜欢列表
	List(c *gin.Context)
}
