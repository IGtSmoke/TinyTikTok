// Package controller contains implementation of interface
package controller

import (
	"TinyTikTok/service/impl"
	"TinyTikTok/utils"
	"errors"
	"github.com/gin-gonic/gin"
	"mime/multipart"
)

var publishServiceImpl = impl.PublishServiceImpl{}

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

	response, err := publishServiceImpl.Action(userId, title, data)
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
	response, err := publishServiceImpl.List(myId, userIdStr)
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
