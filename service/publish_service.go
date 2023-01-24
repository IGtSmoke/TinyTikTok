package service

import "github.com/gin-gonic/gin"

type PublishService interface {
	Action(c *gin.Context)
	List(c *gin.Context)
}
