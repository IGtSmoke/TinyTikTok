package impl

import (
	"TinyTikTok/conf/setup"
	"TinyTikTok/dao"
	"TinyTikTok/model/dto"
	"TinyTikTok/service"
	"TinyTikTok/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v9"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"strconv"
)

type UserServiceImpl struct {
	service.FollowService
}

// Login 用户登录
func (usi *UserServiceImpl) Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	if utils.PasswordInvalid(password) {
		utils.Fail(c, "密码无效")
		return
	}

	user := dao.SearchUserByUserName(username)
	if (user == dto.UserDTO{}) {
		utils.Fail(c, "当前用户不存在")
		return
	}

	if user.Password != password {
		utils.Fail(c, "账号和密码不匹配")
		return
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
		log.Err(err)
	}
	//返回token和用户id
	response := dto.UserLoginResponse{}
	utils.InitSuccessResult(&response.Result)
	response.Token = token
	response.UserID = user.UserID
	utils.Success(c, response)
}

// Register 用户注册
func (usi *UserServiceImpl) Register(c *gin.Context) {
	//接收邮箱和密码
	username := c.Query("username")
	password := c.Query("password")
	//根据username查询用户是否存在
	if utils.PasswordInvalid(password) {
		utils.Fail(c, "密码无效")
		return
	}
	user := dao.SearchUserByUserName(username)
	if (user != dto.UserDTO{}) {
		utils.Fail(c, "用户已存在")
		return
	}

	//保存用户信息到数据库
	userId := setup.SnowflakeNode.Generate().Int64()
	user.UserID = userId
	user.UserName = username
	user.Password = password
	success := dao.SaveUser(&user)
	if !success {
		log.Error().Msg("创建用户失败")
		utils.Fail(c, "创建用户失败")
		return
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
		log.Err(err)
	}

	//返回token和用户id
	response := dto.UserLoginResponse{}
	utils.InitSuccessResult(&response.Result)
	response.Token = token
	response.UserID = userId
	utils.Success(c, response)
	return
}

// UserInfo 查询用户信息
func (usi *UserServiceImpl) UserInfo(c *gin.Context) {
	//获取查询用户的用户信息
	value := c.Query("user_id")
	userId, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		log.Err(err)
	}
	myuserId, exists := c.Get("userId")
	if exists == false {
		log.Error().Msg("[UserInfo]无法得到userId")
		utils.Fail(c, "bad param")
		return
	}
	myId := myuserId.(string)
	parseInt, err := strconv.ParseInt(myId, 10, 64)
	if err != nil {
		log.Err(err)
	}
	userInfo := GetUserInfo(parseInt, userId)
	response := dto.UserInfoResponse{
		User: userInfo,
	}
	utils.InitSuccessResult(&response.Result)
	utils.Success(c, response)
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
		log.Err(err)
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
