package impl

import (
	"TinyTikTok/conf/setup"
	"TinyTikTok/dao"
	"TinyTikTok/model/dto"
	"TinyTikTok/service"
	"TinyTikTok/utils"
	"strconv"
)

type RelationServiceImpl struct {
	service.FollowService
}

// 关注用户
func (i RelationServiceImpl) FollowUser(myId int64, userId int64) (dto.Result, error) {
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
func (i RelationServiceImpl) UnFollowUser(myId int64, userId int64) (dto.Result, error) {
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
