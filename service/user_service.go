package service

import (
	"TinyTikTok/model/dto"
)

// UserService 用户接口
type UserService interface {
	// Login 用户登录
	Login(string, string) (dto.UserLoginResponse, error)
	// Register 用户注册
	Register(string, string) (dto.UserLoginResponse, error)
	// UserInfo 用户信息
	UserInfo(int64, int64) (dto.UserInfoResponse, error)
}
