package utils

import (
	"errors"
	"github.com/gin-gonic/gin"
)

// GetUserIdByMiddleware 通过上下文获取userId
func GetUserIdByMiddleware(c *gin.Context) (int64, error) {
	value := c.GetInt64("userId")
	if value == 0 {
		return 0, errors.New("userId不存在")
	}
	return value, nil
}
