package dao

import (
	"TinyTikTok/conf/setup"
	"TinyTikTok/model/dto"
	"gorm.io/gorm"
)

type LikePO struct {
	gorm.Model
	dto.LikeDTO
}

func (l LikePO) TableName() string {
	return "likes"
}

func QueryLikeByVideoIdAndMyId(myId int64, videoId int64) dto.LikeDTO {
	likeDTO := dto.LikeDTO{}
	setup.Mdb.Model(LikePO{}).Where("user_id = ?", myId, "video_id = ?", videoId).First(&likeDTO)
	return likeDTO
}

func UpdateLike(likeDTO dto.LikeDTO) {
	setup.Mdb.Model(LikePO{}).Updates(&likeDTO)
}

func CreateLike(likeDTO dto.LikeDTO) {
	setup.Mdb.Model(LikePO{}).Create(&likeDTO)
}
