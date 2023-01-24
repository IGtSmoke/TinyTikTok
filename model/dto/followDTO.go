package dto

type Follow struct {
	Id         int64
	UserId     int64
	FollowerId int64
	Cancel     int8
}
