package service

import (
	"TinyTikTok/model/dto"
	"time"
)

type FeedService interface {

	// Feed 视频流
	Feed(time.Time, int64) (dto.PublishFeedResponse, error)
}
