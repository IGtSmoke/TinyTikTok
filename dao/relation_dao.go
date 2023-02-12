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
