// Package controller contains implementation of interface
package controller

import (
	"TinyTikTok/service/impl"
	"TinyTikTok/utils"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"mime/multipart"
	"strconv"
	"time"
)

var psi = impl.PublishServiceImpl{}

// Action 上传视频
func Action(c *gin.Context) {
	//获取相关信息
	userId, err := utils.GetUserIdByMiddleware(c)
	if err != nil {
		utils.FailResponse(c, err)
		return
	}
	title, err := getTitle(c)
	if err != nil {
		utils.FailResponse(c, err)
		return
	}
	data, err := getFile(c)
	if err != nil {
		utils.FailResponse(c, err)
		return
	}

	response, err := psi.Action(userId, title, data)
	if err != nil {
		utils.FailResponse(c, err)
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
		utils.FailResponse(c, err)
		return
	}

	userIdStr := c.Query("user_id")
	response, err := psi.List(myId, userIdStr)
	if err != nil {
		utils.FailResponse(c, err)
		return
	} else {
		utils.Success(c, response)
		return
	}
}

func Feed(c *gin.Context) {
	lastTime, err := getLastTime(c)
	if err != nil {
		log.Err(err)
		lastTime = time.Now()
	}
	myId, _ := utils.GetUserIdByMiddleware(c)
	response, err := psi.Feed(lastTime, myId)
	if err != nil {
		utils.FailResponse(c, err)
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
	if value == "0" {
		lastTime = time.Now()
		return
	}
	inputTime, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		log.Err(err)
		return
	}
	lastTime = time.Unix(inputTime, 0)
	return
}
