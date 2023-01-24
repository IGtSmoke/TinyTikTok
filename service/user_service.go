package service

import "github.com/gin-gonic/gin"

type UserService interface {
	Login(c *gin.Context)
	Register(c *gin.Context)
}
