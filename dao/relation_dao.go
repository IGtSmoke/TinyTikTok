package dao

import (
	"TinyTikTok/conf/setup"
	"TinyTikTok/model/dto"

	"gorm.io/gorm"
)

type RelationPO struct {
	gorm.Model
	dto.FollowDTO
}

func (r RelationPO) TableName() string {
	return "follows"
}

func QueryFollowByMyIdAndUserId(myId int64, userId int64) dto.FollowDTO {
	followDTO := dto.FollowDTO{}
	setup.Mdb.Model(RelationPO{}).Where("user_id = ? AND follower_id = ?", myId, userId).First(&followDTO)
	return followDTO
}

func UpdateFollow(followDTO dto.FollowDTO) {
	setup.Mdb.Model(RelationPO{}).Where("user_id = ? AND follower_id = ?", followDTO.UserID, followDTO.FollowerID).Update("cancel", followDTO.Cancel)
}

func CreateFollow(followDTO dto.FollowDTO) {
	setup.Mdb.Model(RelationPO{}).Create(&RelationPO{
		FollowDTO: followDTO,
	})
}

// 根据用户ID取关注用户ID
func QueryFollowArrayByUserId(userId int64) (followIdArr []int64) {
	setup.Mdb.Model(RelationPO{}).Select("follower_id").Where("user_id = ?", userId).Find(&followIdArr)
	return followIdArr
}

// 根据用户ID取粉丝用户ID
func QueryFollowerArrayByUserId(userId int64) (followerIdArr []int64) {
	setup.Mdb.Model(RelationPO{}).Select("user_id").Where("follower_id = ?", userId).Find(&followerIdArr)
	return followerIdArr
}

// 根据用户ID查询互关用户ID
func QueryFriendArrayByUserId(userId int64) (FriendIdArr []int64) {
	setup.Mdb.Raw("select a.* from (select a.user_id from follows as a inner join follows as bon a.follower_id = ? and b.follower_id = ?) a group by a.user_id;", userId, userId).Find(&FriendIdArr)
	return FriendIdArr
}
