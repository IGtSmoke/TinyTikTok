package service

import "github.com/gin-gonic/gin"

// FavoriteService 赞接口
type FavoriteService interface {
	//Action 赞操作
	Action(c *gin.Context)
	//List 喜欢列表
	List(c *gin.Context)
}
