package utils

import (
	"github.com/rs/zerolog/log"
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

/*
Fail 请求失败
@param errorMsg 返回错误信息
*/
func Fail(c *gin.Context, errorMsg string) {
	c.JSON(http.StatusForbidden, dto.Result{
		StatusCode: -1,
		StatusMsg:  errorMsg,
	})
}

// InitSuccessResult 函数执行成功调用初始化默认结构体
func InitSuccessResult(r *dto.Result) {
	r.StatusCode = 0
	r.StatusMsg = ""
}

func FailResponse(c *gin.Context, err error) {
	log.Err(err)
	Fail(c, err.Error())
	return
}
