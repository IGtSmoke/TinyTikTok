package service

import "github.com/gin-gonic/gin"

// UserService 用户接口
type UserService interface {
	// Login 用户登录
	Login(c *gin.Context)
	// Register 用户注册
	Register(c *gin.Context)
}
