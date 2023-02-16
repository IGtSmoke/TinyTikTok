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

func GetVideoByVideoId(videoId int64) dto.VideoDTO {
	var videoDTO dto.VideoDTO
	setup.Mdb.Model(&VideoPO{}).Where("id = ?", videoId).Order("created_at DESC").First(&videoDTO)
	return videoDTO
}

func GetVideosAndNextTimeByLastTime(lastTime time.Time) ([]dto.VideoDTO, time.Time) {
	var videoDTOS []dto.VideoDTO
	//maxTime ：9999-12-31 23:59:59
	const maxTime = 253402300799
	timestamp := lastTime.UnixNano()
	if timestamp > maxTime || timestamp < 0 {
		lastTime = time.Unix(150000000000, 0)
	}

	setup.Mdb.Model(&VideoPO{}).Where("updated_at < ?", lastTime).Order("updated_at desc").Limit(30).Find(&videoDTOS)
	var nextTime time.Time
	if len(videoDTOS) == 0 {
		return videoDTOS, lastTime
	}
	setup.Mdb.Model(&VideoPO{}).Select("updated_at").Where("id = ?", videoDTOS[len(videoDTOS)-1].Id).Find(&nextTime)
	return videoDTOS, nextTime
}
