package dao

import (
	"TinyTikTok/conf/setup"
	"TinyTikTok/model/dto"
	"gorm.io/gorm"
	"time"
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

func GetVideosByAuthorId(authorId int64) []dto.VideoDTO {
	var videoDTOS []dto.VideoDTO
	setup.Mdb.Model(&VideoPO{}).Where("author_id = ?", authorId).Order("created_at DESC").Find(&videoDTOS)
	return videoDTOS
}

func GetVideosAndNextTimeByLastTime(lastTime time.Time) ([]dto.VideoDTO, time.Time) {
	var videoDTOS []dto.VideoDTO
	//todo 传入时间有误 year is not in the range [1, 9999]: 55065
	setup.Mdb.Model(&VideoPO{}).Where("updated_at < ?", lastTime).Order("updated_at desc").Limit(30).Find(&videoDTOS)
	var nextTime time.Time
	if len(videoDTOS) == 0 {
		return videoDTOS, lastTime
	}
	setup.Mdb.Model(&VideoPO{}).Select("updated_at").Where("id = ?", videoDTOS[len(videoDTOS)-1].Id).Find(&nextTime)
	return videoDTOS, nextTime
}
