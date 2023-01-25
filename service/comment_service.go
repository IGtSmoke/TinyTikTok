package service

import "github.com/gin-gonic/gin"

// CommentService 评论接口
type CommentService interface {
	//Action 评论操作
	Action(c *gin.Context)
	//List 评论列表
	List(c *gin.Context)
}
