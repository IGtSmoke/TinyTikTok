package controller

import (
	"TinyTikTok/conf/setup"
	"TinyTikTok/service/impl"
	"TinyTikTok/utils"
	"runtime/debug"
	"strconv"

	"github.com/gin-gonic/gin"
)

var commentServiceImpl = impl.CommentServiceImpl{}

// Comment 评论
func Comment(c *gin.Context) {
	videoId, err := getVideoId(c)
	if err != nil {
		utils.Fail(c, err)
		return
	}
	actionType, err := getActionType(c)
	if err != nil {
		utils.Fail(c, err)
		return
	}

	myId, err := utils.GetUserIdByMiddleware(c)
	if err != nil {
		utils.Fail(c, err)
		return
	}

	if actionType == 1 {
		response, err := commentServiceImpl.Comment(myId, videoId, c.Query("comment_text"))
		if err != nil {
			utils.Fail(c, err)
			return
		} else {
			utils.Success(c, response)
			return
		}
	} else if actionType == 2 {
		commentId, err := strconv.ParseInt(c.Query("comment_id"), 10, 64)
		if err != nil {
			setup.Logger("common").Err(err).Interface("stack", string(debug.Stack())).Send()
		}
		response, err := commentServiceImpl.DeleteComment(commentId)
		if err != nil {
			utils.Fail(c, err)
			return
		} else {
			utils.Success(c, response)
			return
		}
	}
}

func getActionType(c *gin.Context) (int8, error) {
	value := c.Query("action_type")
	if value == "1" {
		return 1, nil
	} else {
		return 2, nil
	}
}

func CommentList(c *gin.Context) {
	myId, _ := utils.GetUserIdByMiddleware(c)
	videoId, err := strconv.ParseInt(c.Query("video_id"), 10, 64)
	if err != nil {
		setup.Logger("common").Err(err).Interface("stack", string(debug.Stack())).Send()
	}
	response, err := commentServiceImpl.List(myId, videoId)
	if err != nil {
		utils.Fail(c, err)
		return
	} else {
		utils.Success(c, response)
		return
	}
}
