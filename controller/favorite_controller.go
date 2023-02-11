package controller

import (
	"TinyTikTok/service/impl"
	"TinyTikTok/utils"
	"github.com/gin-gonic/gin"
	"strconv"
)

var favoriteServiceImpl = impl.FavoriteServiceImpl{}

// Thumb 点赞
func Thumb(c *gin.Context) {
	videoId, err := getVideoId(c)
	if err != nil {
		utils.Fail(c, err)
		return
	}
	isThumb, err := getIsThumb(c)
	myId, err := utils.GetUserIdByMiddleware(c)
	if err != nil {
		utils.Fail(c, err)
		return
	}
	response, err := favoriteServiceImpl.Thumb(myId, videoId, isThumb)
	if err != nil {
		utils.Fail(c, err)
		return
	} else {
		utils.Success(c, response)
		return
	}
}

func getVideoId(c *gin.Context) (int64, error) {
	value := c.Query("video_id")
	videoId, err := strconv.ParseInt(value, 10, 64)
	return videoId, err
}

func getIsThumb(c *gin.Context) (bool, error) {
	value := c.Query("action_type")
	if value == "1" {
		return true, nil
	} else {
		return false, nil
	}
}
