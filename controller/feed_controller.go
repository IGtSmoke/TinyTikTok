package controller

import (
	"TinyTikTok/conf/setup"
	"TinyTikTok/service/impl"
	"TinyTikTok/utils"
	"github.com/gin-gonic/gin"
	"runtime/debug"
	"strconv"
	"time"
)

var feedServiceImpl = impl.FeedServiceImpl{}

func Feed(c *gin.Context) {
	lastTime, err := getLastTime(c)
	if err != nil {
		setup.Logger("common").Err(err).Send()
		lastTime = time.Now()
	}
	myId, _ := utils.GetUserIdByMiddleware(c)
	response, err := feedServiceImpl.Feed(lastTime, myId)
	if err != nil {
		utils.Fail(c, err)
		return
	} else {
		utils.Success(c, response)
		return
	}
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
