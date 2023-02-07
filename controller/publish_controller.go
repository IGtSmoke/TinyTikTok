// Package controller contains implementation of interface
package controller

import (
	"TinyTikTok/conf/setup"
	"TinyTikTok/service/impl"
	"TinyTikTok/utils"
	"errors"
	"github.com/gin-gonic/gin"
	"mime/multipart"
	"runtime/debug"
	"strconv"
	"time"
)

var psi = impl.PublishServiceImpl{}

// Action 上传视频
func Action(c *gin.Context) {
	//获取相关信息
	userId, err := utils.GetUserIdByMiddleware(c)
	if err != nil {
		utils.Fail(c, err)
		return
	}
	title, err := getTitle(c)
	if err != nil {
		utils.Fail(c, err)
		return
	}
	data, err := getFile(c)
	if err != nil {
		utils.Fail(c, err)
		return
	}

	response, err := psi.Action(userId, title, data)
	if err != nil {
		utils.Fail(c, err)
		return
	} else {
		utils.Success(c, response)
		return
	}
}

// List 视频列表
func List(c *gin.Context) {
	myId, err := utils.GetUserIdByMiddleware(c)
	if err != nil {
		utils.Fail(c, err)
		return
	}

	userIdStr := c.Query("user_id")
	response, err := psi.List(myId, userIdStr)
	if err != nil {
		utils.Fail(c, err)
		return
	} else {
		utils.Success(c, response)
		return
	}
}

func Feed(c *gin.Context) {
	lastTime, err := getLastTime(c)
	if err != nil {
		setup.Logger("common").Err(err).Send()
		lastTime = time.Now()
	}
	myId, _ := utils.GetUserIdByMiddleware(c)
	response, err := psi.Feed(lastTime, myId)
	if err != nil {
		utils.Fail(c, err)
		return
	} else {
		utils.Success(c, response)
		return
	}
}

func getTitle(c *gin.Context) (string, error) {
	title := c.PostForm("title")
	if title == "" {
		return "", errors.New("title不能为空")
	}
	return title, nil
}

func getFile(c *gin.Context) (*multipart.FileHeader, error) {
	data, err := c.FormFile("data")
	if err != nil {
		return nil, err
	}
	return data, nil
}

func getLastTime(c *gin.Context) (lastTime time.Time, err error) {
	value := c.Query("latest_time")
	if value == "" {
		lastTime = time.Now()
		return
	}
	inputTime, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		setup.Logger("common").Err(err).Interface("stack", string(debug.Stack())).Send()
		return
	}
	lastTime = time.Unix(inputTime, 0)
	return
}
