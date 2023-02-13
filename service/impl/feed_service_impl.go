package impl

import (
	"TinyTikTok/dao"
	"TinyTikTok/model/dto"
	"TinyTikTok/utils"
	"time"
)

type FeedServiceImpl struct {
}

func (f FeedServiceImpl) Feed(lastTime time.Time, myId int64) (dto.PublishFeedResponse, error) {

	result := make([]dto.Video, 0, 30)
	videoDTOS, timestamp := dao.GetVideosAndNextTimeByLastTime(lastTime)
	for _, videoDTO := range videoDTOS {
		assembleUser(&result, myId, videoDTO)
	}
	nextTime := timestamp.Unix()
	response := dto.PublishFeedResponse{
		NextTime:  &nextTime,
		VideoList: result,
	}
	utils.InitSuccessResult(&response.Result)
	return response, nil
}
