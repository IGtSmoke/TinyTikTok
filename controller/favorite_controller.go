package controller

import (
	"TinyTikTok/service/impl"
	"TinyTikTok/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

var fsi = impl.FavoriteServiceImpl{}

// Thumb 点赞
func Thumb(c *gin.Context) {
	videoId, err := getVideoId(c)
	if err != nil {
		utils.Fail(c, err)
		return
	}
	isThumb, err := getIsThumb(c)
	if err != nil {
		utils.Fail(c, err)
		return
	}
	myId, err := utils.GetUserIdByMiddleware(c)
	if err != nil {
		utils.Fail(c, err)
		return
	}
	response, err := fsi.Thumb(myId, videoId, isThumb)
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

func getIsThumb(c *gin.Context) (int8, error) {
	value := c.Query("action_type")
	if value == "1" {
		return 1, nil
	} else {
		return 0, nil
	}
}
