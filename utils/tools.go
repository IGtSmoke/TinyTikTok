package utils

import (
	"encoding/base64"
	"errors"
	"github.com/gin-gonic/gin"
)

// Base64Encode 加密 []byte -> string
func Base64Encode(src []byte) string {
	result := base64.StdEncoding.EncodeToString(src)

	return result
}

// Base64Decode 解密 string -> []byte
func Base64Decode(src string) ([]byte, error) {
	result, err := base64.StdEncoding.DecodeString(src)

	return result, err
}

// GetUserIdByMiddleware 通过上下文获取userId
func GetUserIdByMiddleware(c *gin.Context) (int64, error) {
	value := c.GetInt64("userId")
	if value == 0 {
		return 0, errors.New("userId不存在")
	}
	return value, nil
}
