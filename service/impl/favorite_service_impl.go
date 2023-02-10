package impl

import (
	"TinyTikTok/conf/setup"
	"TinyTikTok/dao"
	"TinyTikTok/model/dto"
	"TinyTikTok/utils"
	"strconv"
)

type FavoriteServiceImpl struct {
}

func (i FavoriteServiceImpl) Thumb(myId int64, videoId int64, isThumb bool) (dto.Result, error) {
	likeDTO := dao.QueryLikeByVideoIdAndMyId(myId, videoId)
	if (likeDTO == dto.LikeDTO{}) {
		dao.CreateLike(dto.LikeDTO{
			UserId:  myId,
			VideoId: videoId,
			IsThumb: isThumb,
		})
	} else {
		dao.UpdateLike(dto.LikeDTO{
			UserId:  myId,
			VideoId: videoId,
			IsThumb: isThumb,
		})
	}
	favoriteKey := utils.VideoLikeKey + strconv.FormatInt(videoId, 10)
	if isThumb {
		setup.Rdb.SAdd(setup.Rctx, favoriteKey, myId)
	} else {
		setup.Rdb.SRem(setup.Rctx, favoriteKey, myId)
	}
	response := dto.Result{}
	utils.InitSuccessResult(&response)
	return response, nil
}
