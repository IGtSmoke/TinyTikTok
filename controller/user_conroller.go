package controller

import (
	"TinyTikTok/service/impl"
	"github.com/gin-gonic/gin"
)

var usi = impl.UserServiceImpl{}

// Login 用户登录
func Login(c *gin.Context) {
	usi.Login(c)
}

// Register 用户注册
func Register(c *gin.Context) {
	usi.Register(c)
}

func UserInfo(c *gin.Context) {
	usi.UserInfo(c)
}
