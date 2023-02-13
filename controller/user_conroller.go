package controller

import (
	"TinyTikTok/service/impl"
	"TinyTikTok/utils"
	"errors"
	"github.com/duke-git/lancet/v2/validator"
	"github.com/gin-gonic/gin"
	"strconv"
)

var userServiceImpl = impl.UserServiceImpl{}

// Login 用户登录
func Login(c *gin.Context) {
	//获取用户名和密码
	username, password, err := getUsernameAndPassword(c)
	if err != nil {
		utils.Fail(c, err)
		return
	}

	response, err := userServiceImpl.Login(username, password)

	if err != nil {
		utils.Fail(c, err)
		return
	} else {
		utils.Success(c, response)
		return
	}
}

// Register 用户注册
func Register(c *gin.Context) {
	username, password, err := getUsernameAndPassword(c)
	if err != nil {
		utils.Fail(c, err)
		return
	}

	response, err := userServiceImpl.Register(username, password)

	if err != nil {
		utils.Fail(c, err)
		return
	} else {
		utils.Success(c, response)
		return
	}
}

func UserInfo(c *gin.Context) {
	value := c.Query("user_id")
	userId, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		utils.Fail(c, err)
		return
	}
	myId, err := utils.GetUserIdByMiddleware(c)
	if err != nil {
		utils.Fail(c, err)
		return
	}

	response, err := userServiceImpl.UserInfo(myId, userId)

	if err != nil {
		utils.Fail(c, err)
		return
	} else {
		utils.Success(c, response)
		return
	}
}

func getUsernameAndPassword(c *gin.Context) (username string, password string, err error) {
	username = c.Query("username")
	password = c.Query("password")

	if validator.IsEmptyString(username) || utils.PasswordInvalid(password) {
		return "", "", errors.New("用户名或密码格式错误")
	}
	return
}
