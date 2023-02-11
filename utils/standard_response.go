package utils

import (
	"TinyTikTok/conf/setup"
	"net/http"

	"TinyTikTok/model/dto"
	"github.com/gin-gonic/gin"
)

/*
Success 请求成功
@param data 返回相应接口的结构体
*/
func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, data)
}

// InitSuccessResult 函数执行成功调用初始化默认结构体
func InitSuccessResult(r *dto.Result) {
	r.StatusCode = 0
	r.StatusMsg = ""
}

/*
Fail 请求失败
@param errorMsg 返回错误信息
*/
func Fail(c *gin.Context, err error) {
	setup.Logger("common").Err(err).Send()
	c.JSON(http.StatusForbidden, dto.Result{
		StatusCode: -1,
		StatusMsg:  err.Error(),
	})
	return
}
