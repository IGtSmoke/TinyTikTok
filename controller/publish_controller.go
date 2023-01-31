// Package controller contains implementation of interface
package controller

import (
	"TinyTikTok/service/impl"
	"github.com/gin-gonic/gin"
)

var psi = impl.PublishServiceImpl{}

// Action 上传视频
func Action(c *gin.Context) {
	psi.Action(c)
}

// List 视频列表
func List(c *gin.Context) {
	psi.List(c)
}

func Feed(c *gin.Context) {
	psi.Feed(c)
}
