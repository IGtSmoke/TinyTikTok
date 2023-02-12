package controller

import (
	"TinyTikTok/service/impl"
	"TinyTikTok/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

var rsi = impl.RelationServiceImpl{}

// 关注操作
func Follow(c *gin.Context) {
	myId, err := utils.GetUserIdByMiddleware(c)
	if err != nil {
		utils.Fail(c, err)
		return
	}
	// token := c.Query("token")
	// tokenKey := utils.LoginUserKey + token
	// // 取出userId
	// tmp, err := setup.Rdb.HGet(setup.Rctx, tokenKey, "userId").Result()
	// if err != nil {
	// 	utils.Fail(c, err)
	// 	return
	// }
	// myId, err := strconv.ParseInt(tmp, 10, 64)
	// if err != nil {
	// 	utils.Fail(c, err)
	// 	return
	// }
	//关注目标id
	value := c.Query("to_user_id")
	user_id, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		utils.Fail(c, err)
		return
	}
	//1-关注，2-取消关注
	action_type, err := getIsFollow(c)
	if err != nil {
		utils.Fail(c, err)
		return
	}
	if action_type == 1 {
		response, err := rsi.FollowUser(myId, user_id)
		if err != nil {
			utils.Fail(c, err)
			return
		} else {
			utils.Success(c, response)
			return
		}
	} else {
		response, err := rsi.UnFollowUser(myId, user_id)
		if err != nil {
			utils.Fail(c, err)
			return
		} else {
			utils.Success(c, response)
			return
		}
	}
}

// 关注列表
func FollowList(c *gin.Context) {

}

// 粉丝列表
func FollowerList(c *gin.Context) {

}

func getIsFollow(c *gin.Context) (int8, error) {
	value := c.Query("action_type")
	if value == "1" {
		return 1, nil
	} else {
		return 0, nil
	}
}
