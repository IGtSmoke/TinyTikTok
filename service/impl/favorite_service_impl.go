package impl

import (
	"TinyTikTok/conf/setup"
	"TinyTikTok/dao"
	"TinyTikTok/model/dto"
	"TinyTikTok/utils"
	"runtime/debug"
	"strconv"
)

type FavoriteServiceImpl struct {
}

func (i FavoriteServiceImpl) Thumb(myId int64, videoId int64, isThumb int8) (dto.Result, error) {
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
	if isThumb == 1 {
		setup.Rdb.SAdd(setup.Rctx, favoriteKey, myId)
	} else {
		setup.Rdb.SRem(setup.Rctx, favoriteKey, myId)
	}
	response := dto.Result{}
	utils.InitSuccessResult(&response)
	return response, nil
}

func (i FavoriteServiceImpl) List(myId int64, userIdStr string) (dto.FavoriteListResponse, error) {
	response := dto.FavoriteListResponse{}
	result := make([]dto.Video, 0)
	authorId, err := strconv.ParseInt(userIdStr, 10, 64)
	if err != nil {
		setup.Logger("common").Err(err).Interface("stack", string(debug.Stack())).Send()
	}
	likeDTOS := dao.QueryLikeByUserId(authorId)
	for _, tmp := range likeDTOS {
		assembleVideoAndUser(&result, authorId, dao.GetVideoByVideoId(tmp.VideoId))
	}
	response.VideoList = result
	utils.InitSuccessResult(&response.Result)
	return response, nil
}
