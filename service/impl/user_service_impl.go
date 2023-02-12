package impl

import (
	"TinyTikTok/conf/setup"
	"TinyTikTok/dao"
	"TinyTikTok/model/dto"
	"TinyTikTok/service"
	"TinyTikTok/utils"
	"errors"
	"strconv"

	"github.com/go-redis/redis/v9"
	"github.com/google/uuid"
)

type UserServiceImpl struct {
	service.FollowService
}

// Login 用户登录
func (usi *UserServiceImpl) Login(username string, password string) (dto.UserLoginResponse, error) {
	response := dto.UserLoginResponse{}
	user := dao.SearchUserByUserName(username)
	if user == (dto.UserDTO{}) {
		return response, errors.New("用户不存在")
	}

	if user.Password != password {
		return response, errors.New("密码错误")
	}
	//登录成功
	//生成token
	token := uuid.NewString()
	tokenKey := utils.LoginUserKey + token
	//通过管道包装发送到redis的用户信息
	if _, err := setup.Rdb.Pipelined(setup.Rctx, func(rdb redis.Pipeliner) error {
		rdb.HSet(setup.Rctx, tokenKey, "userId", user.UserID)
		rdb.Expire(setup.Rctx, tokenKey, utils.LoginUserTTL)
		return nil
	}); err != nil {
		setup.Logger("common").Err(err).Send()
	}
	//返回token和用户id

	utils.InitSuccessResult(&response.Result)
	response.Token = token
	response.UserID = user.UserID
	return response, nil
}

// Register 用户注册
func (usi *UserServiceImpl) Register(username string, password string) (dto.UserLoginResponse, error) {
	response := dto.UserLoginResponse{}
	user := dao.SearchUserByUserName(username)
	if user != (dto.UserDTO{}) {
		return response, errors.New("用户已存在")
	}

	//保存用户信息到数据库
	userId := setup.SnowflakeNode.Generate().Int64()
	user.UserID = userId
	user.UserName = username
	user.Password = password
	success := dao.SaveUser(&user)
	if !success {
		return response, errors.New("保存用户失败")
	}

	//保存用户信息到redis
	//生成token
	token := uuid.NewString()
	tokenKey := utils.LoginUserKey + token
	//通过管道包装发送到redis的用户信息
	if _, err := setup.Rdb.Pipelined(setup.Rctx, func(rdb redis.Pipeliner) error {
		rdb.HSet(setup.Rctx, tokenKey, "userId", user.UserID)
		rdb.Expire(setup.Rctx, tokenKey, utils.LoginUserTTL)
		return nil
	}); err != nil {
		setup.Logger("common").Err(err).Send()
	}

	//返回token和用户id
	utils.InitSuccessResult(&response.Result)
	response.Token = token
	response.UserID = userId
	return response, nil
}

// UserInfo 查询用户信息
func (usi *UserServiceImpl) UserInfo(myId int64, userId int64) (dto.UserInfoResponse, error) {
	//获取查询用户的用户信息
	userInfo := GetUserInfo(myId, userId)
	response := dto.UserInfoResponse{
		User: userInfo,
	}
	utils.InitSuccessResult(&response.Result)
	return response, nil
}

// GetUserInfo 查询对方用户信息及自己是否关注对方
func GetUserInfo(myId int64, userId int64) dto.User {
	userDTO := dao.SearchUserByUserId(userId)
	//获取自己的用户信息

	//根据userId查询redis获得关注数和粉丝数
	fansKey := utils.FansUserKey + strconv.FormatInt(userDTO.UserID, 10)
	followKey := utils.FollowUserKey + strconv.FormatInt(userDTO.UserID, 10)
	//构建管道
	pipeline := setup.Rdb.Pipeline()
	pipeline.SCard(setup.Rctx, fansKey)
	pipeline.SCard(setup.Rctx, followKey)
	pipeline.SIsMember(setup.Rctx, fansKey, myId)
	exec, err := pipeline.Exec(setup.Rctx)
	if err != nil {
		setup.Logger("common").Err(err).Send()
	}
	//获得结果
	fansCount := exec[0].(*redis.IntCmd).Val()
	followCount := exec[1].(*redis.IntCmd).Val()
	isFollow := exec[2].(*redis.BoolCmd).Val()

	user := dto.User{
		FollowCount:   followCount,
		FollowerCount: fansCount,
		ID:            userDTO.UserID,
		IsFollow:      isFollow,
		Name:          userDTO.UserName,
	}
	return user
}
