package impl

import (
	"TinyTikTok/conf/setup"
	"TinyTikTok/dao"
	"TinyTikTok/model/dto"
	"TinyTikTok/service"
	"TinyTikTok/utils"
	"strconv"
)

type FollowServiceImpl struct {
	service.FollowService
}

type RelationServiceImpl struct {
	service.RelationService
}

// 关注用户
func (i FollowServiceImpl) FollowUser(myId int64, userId int64) (dto.Result, error) {
	followDTO := dao.QueryFollowByMyIdAndUserId(myId, userId)
	if (followDTO == dto.FollowDTO{}) {
		dao.CreateFollow(dto.FollowDTO{
			UserID:     myId,
			FollowerID: userId,
			Cancel:     0,
		})
	} else {
		dao.UpdateFollow(dto.FollowDTO{
			UserID:     myId,
			FollowerID: userId,
			Cancel:     0,
		})
	}
	fansKey := utils.FansUserKey + strconv.FormatInt(userId, 10)
	followKey := utils.FollowUserKey + strconv.FormatInt(myId, 10)
	setup.Rdb.SAdd(setup.Rctx, fansKey, myId)
	setup.Rdb.SAdd(setup.Rctx, followKey, userId)
	// if _, err := setup.Rdb.Pipelined(setup.Rctx, func(rdb redis.Pipeliner) error {
	// 	rdb.SAdd(setup.Rctx, fansKey, myId)
	// 	rdb.SAdd(setup.Rctx, followKey, userId)
	// 	return nil
	// }); err != nil {
	// 	setup.Logger("common").Err(err).Send()
	// }
	response := dto.Result{}
	utils.InitSuccessResult(&response)
	return response, nil
}

// 取关用户
func (i FollowServiceImpl) UnFollowUser(myId int64, userId int64) (dto.Result, error) {
	followDTO := dao.QueryFollowByMyIdAndUserId(myId, userId)
	if (followDTO == dto.FollowDTO{}) {
		dao.CreateFollow(dto.FollowDTO{
			UserID:     myId,
			FollowerID: userId,
			Cancel:     1,
		})
	} else {
		dao.UpdateFollow(dto.FollowDTO{
			UserID:     myId,
			FollowerID: userId,
			Cancel:     1,
		})
	}
	fansKey := utils.FansUserKey + strconv.FormatInt(userId, 10)
	followKey := utils.FollowUserKey + strconv.FormatInt(myId, 10)
	setup.Rdb.SRem(setup.Rctx, fansKey, myId)
	setup.Rdb.SRem(setup.Rctx, followKey, userId)
	// if _, err := setup.Rdb.Pipelined(setup.Rctx, func(rdb redis.Pipeliner) error {
	// 	rdb.SRem(setup.Rctx, fansKey, myId)
	// 	rdb.SRem(setup.Rctx, followKey, userId)
	// 	return nil
	// }); err != nil {
	// 	setup.Logger("common").Err(err).Send()
	// }
	response := dto.Result{}
	utils.InitSuccessResult(&response)
	return response, nil
}

// fansKey := utils.FansUserKey + strconv.FormatInt(myId, 10)
// followKey := utils.FollowUserKey + strconv.FormatInt(userDTO.UserID, 10)

// 查看目标关注列表
func (i RelationServiceImpl) ShowFollowList(myId int64, userId int64) (dto.RelationList, error) {
	//从redis中根据UserId取关注用户Id数组
	followKey := utils.FollowUserKey + strconv.FormatInt(userId, 10)
	tmp := setup.Rdb.SMembers(setup.Rctx, followKey).Val()
	var followUserIdArr []int64
	response := dto.RelationList{}
	//没有的话从数据库里找
	if len(tmp) == 0 {
		followUserIdArr = dao.QueryFollowArrayByUserId(myId)
	} else {
		for _, value := range tmp {
			val, _ := strconv.ParseInt(value, 10, 64)
			followUserIdArr = append(followUserIdArr, val)
		}
	}
	followUserInfoArr := []dto.User{}
	for _, value := range followUserIdArr {
		followUserInfoArr = append(followUserInfoArr, GetUserInfo(myId, value))
	}
	response.UserList = followUserInfoArr
	utils.InitSuccessResult(&response.Result)
	return response, nil
}

// 查看目标粉丝列表
func (i RelationServiceImpl) ShowFollowerList(myId int64, userId int64) (dto.RelationList, error) {
	//从redis中根据UserId取粉丝用户Id数组
	fanKey := utils.FansUserKey + strconv.FormatInt(userId, 10)
	tmp := setup.Rdb.SMembers(setup.Rctx, fanKey).Val()
	var followerUserIdArr []int64
	response := dto.RelationList{}
	//没有的话从数据库里找
	if len(tmp) == 0 {
		followerUserIdArr = dao.QueryFollowerArrayByUserId(myId)
	} else {
		for _, value := range tmp {
			val, _ := strconv.ParseInt(value, 10, 64)
			followerUserIdArr = append(followerUserIdArr, val)
		}
	}
	followerUserInfoArr := []dto.User{}
	for _, value := range followerUserIdArr {
		followerUserInfoArr = append(followerUserInfoArr, GetUserInfo(myId, value))
	}
	response.UserList = followerUserInfoArr
	utils.InitSuccessResult(&response.Result)
	return response, nil
}

// 查看目标好友列表
func (i RelationServiceImpl) ShowFriendList(myId int64, userId int64) (dto.RelationList, error) {
	//从redis中根据UserId取粉丝用户Id数组
	fanKey := utils.FansUserKey + strconv.FormatInt(userId, 10)
	fantmp := setup.Rdb.SMembers(setup.Rctx, fanKey).Val()
	followKey := utils.FollowUserKey + strconv.FormatInt(userId, 10)
	followtmp := setup.Rdb.SMembers(setup.Rctx, followKey).Val()
	var tmp []string
	for _, fanvalue := range fantmp {
		for _, followvalue := range followtmp {
			if fanvalue == followvalue {
				tmp = append(tmp, fanvalue)
			}
		}
	}
	var friendUserIdArr []int64
	response := dto.RelationList{}
	//没有的话从数据库里找
	if len(tmp) == 0 {
		friendUserIdArr = dao.QueryFriendArrayByUserId(myId)
	} else {
		for _, value := range tmp {
			val, _ := strconv.ParseInt(value, 10, 64)
			friendUserIdArr = append(friendUserIdArr, val)
		}
	}
	friendUserInfoArr := []dto.User{}
	for _, value := range friendUserIdArr {
		friendUserInfoArr = append(friendUserInfoArr, GetUserInfo(myId, value))
	}
	response.UserList = friendUserInfoArr
	utils.InitSuccessResult(&response.Result)
	return response, nil
}
