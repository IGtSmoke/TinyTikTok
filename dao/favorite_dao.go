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
	setup.Mdb.Model(LikePO{}).Where("user_id = ? AND video_id = ?", myId, videoId).First(&likeDTO)
	return likeDTO
}

func QueryLikeByUserId(userId int64) []dto.LikeDTO {
	var Result []dto.LikeDTO
	setup.Mdb.Model(LikePO{}).Where("user_id = ? AND status = ?", userId, 1).Find(&Result)
	return Result
}

func UpdateLike(likeDTO dto.LikeDTO) {
	setup.Mdb.Model(LikePO{}).Where("user_id = ? AND video_id = ?", likeDTO.UserId, likeDTO.VideoId).Update("status", likeDTO.IsThumb)
}

func CreateLike(likeDTO dto.LikeDTO) {
	setup.Mdb.Create(&LikePO{
		LikeDTO: likeDTO,
	})
}
