package dao

import (
	"TinyTikTok/conf/setup"
	"TinyTikTok/model/dto"
	"gorm.io/gorm"
)

type VideoPO struct {
	gorm.Model
	dto.VideoDTO
}

func (v VideoPO) TableName() string {
	return "videos"
}

func SaveVideo(videoDTO dto.VideoDTO) {
	setup.Mdb.Create(&VideoPO{
		VideoDTO: videoDTO,
	})
}
