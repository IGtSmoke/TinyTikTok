package controller

import (
	"TinyTikTok/service/impl"
	"TinyTikTok/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

var fsi_ = impl.FollowServiceImpl{}
var rsi = impl.RelationServiceImpl{}

// Follow 关注操作
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
	userId, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		utils.Fail(c, err)
		return
	}
	//1-关注，2-取消关注
	actionType, err := getIsFollow(c)
	if err != nil {
		utils.Fail(c, err)
		return
	}
	if actionType == 1 {
		response, err := fsi_.FollowUser(myId, userId)
		if err != nil {
			utils.Fail(c, err)
			return
		} else {
			utils.Success(c, response)
			return
		}
	} else {
		response, err := fsi_.UnFollowUser(myId, userId)
		if err != nil {
			utils.Fail(c, err)
			return
		} else {
			utils.Success(c, response)
			return
		}
	}
}

// FollowList 关注列表
func FollowList(c *gin.Context) {
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
	response, err := rsi.ShowFollowList(myId, userId)
	if err != nil {
		utils.Fail(c, err)
		return
	} else {
		utils.Success(c, response)
		return
	}
}

// FollowerList 粉丝列表
func FollowerList(c *gin.Context) {
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
	response, err := rsi.ShowFollowerList(myId, userId)
	if err != nil {
		utils.Fail(c, err)
		return
	} else {
		utils.Success(c, response)
		return
	}
}

func FriendList(c *gin.Context) {
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
	response, err := rsi.ShowFriendList(myId, userId)
	if err != nil {
		utils.Fail(c, err)
		return
	} else {
		utils.Success(c, response)
		return
	}
}
func getIsFollow(c *gin.Context) (int8, error) {
	value := c.Query("action_type")
	if value == "1" {
		return 1, nil
	} else {
		return 0, nil
	}
}
