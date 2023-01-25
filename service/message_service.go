package service

import "github.com/gin-gonic/gin"

// MessageService 消息接口
type MessageService interface {
	// Chat 聊天记录
	Chat(c *gin.Context)
	// Action 消息操作
	Action(c *gin.Context)
}
