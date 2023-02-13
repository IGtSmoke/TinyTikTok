package impl

import (
	"TinyTikTok/conf/setup"
	"TinyTikTok/dao"
	"TinyTikTok/model/dto"
	"TinyTikTok/utils"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v9"
	"runtime/debug"
	"strconv"
	"time"
)

type CommentServiceImpl struct {
}

func (c CommentServiceImpl) Comment(myId int64, videoId int64, commentText string) (dto.CommentActionResponse, error) {
	commentDTO := dto.CommentDTO{
		UserID:      myId,
		VideoID:     videoId,
		CommentText: commentText,
		CreateDate:  fmt.Sprintf("%02d-%02d", time.Now().Month(), time.Now().Day()),
	}
	commentID := dao.CreateComment(commentDTO)
	commentDTO.Id = commentID

	response := dto.CommentActionResponse{
		Comment: dto.Comment{
			CommentDTO: commentDTO,
			User:       GetUserInfo(myId, myId),
		},
	}
	commentKey := utils.VideoCommentKey + strconv.FormatInt(videoId, 10)

	commentDTOJson, err := json.Marshal(commentDTO)
	if err != nil {
		return response, errors.New("commentDTO json化失败")
	}
	//保存视频信息到redis
	if _, err := setup.Rdb.Pipelined(setup.Rctx, func(rdb redis.Pipeliner) error {
		rdb.LPush(setup.Rctx, commentKey, commentDTOJson)
		return nil
	}); err != nil {
		return response, errors.New("保存comment到redis失败")
	}
	utils.InitSuccessResult(&response.Result)
	return response, nil
}

func (c CommentServiceImpl) DeleteComment(commentId int64) (dto.CommentActionResponse, error) {
	commentDTO := dao.QueryCommentById(commentId)
	if commentDTO == (dto.CommentDTO{}) {
		return dto.CommentActionResponse{}, errors.New("comment not found")
	}
	dao.DeleteCommentById(commentId)
	commentKey := utils.VideoCommentKey + strconv.FormatInt(int64(commentDTO.VideoID), 10)
	//从redis中删除视频信息
	if _, err := setup.Rdb.Pipelined(setup.Rctx, func(rdb redis.Pipeliner) error {
		rdb.LRem(setup.Rctx, commentKey, 0, commentDTO)
		return nil
	}); err != nil {
		return dto.CommentActionResponse{}, errors.New("从redis删除comment失败")
	}
	response := dto.CommentActionResponse{}
	utils.InitSuccessResult(&response.Result)
	return response, nil
}

func (c CommentServiceImpl) List(myId int64, videoId int64) (dto.CommentListResponse, error) {
	response := dto.CommentListResponse{}
	//从redis中查询作者视频信息
	commentKey := utils.VideoCommentKey + strconv.FormatInt(videoId, 10)
	val, err := setup.Rdb.LRange(setup.Rctx, commentKey, 0, -1).Result()
	result := make([]dto.Comment, 0, len(val))
	// 如果从redis中获取成功，则直接返回结果
	if err != nil {
		var commentDTO dto.CommentDTO
		for _, tmp := range val {
			//获取commentDTO
			if err := json.Unmarshal([]byte(tmp), &commentDTO); err != nil {
				setup.Logger("common").Err(err).Interface("stack", string(debug.Stack())).Send()
			}
			assembleCommentAndUser(&result, myId, commentDTO.UserID, commentDTO)
		}
		response.CommentList = result
		utils.InitSuccessResult(&response.Result)
		return response, nil
	}

	//redis查询失败尝试mysql查询
	setup.Logger("common").Err(err).Msg("redis查询失败尝试mysql查询")
	setup.Logger("common").Err(err).Send()
	commentDTOS := dao.QueryVideoComment(videoId)
	for _, tmp := range commentDTOS {
		assembleCommentAndUser(&result, myId, tmp.UserID, tmp)
	}
	response.CommentList = result
	utils.InitSuccessResult(&response.Result)
	return response, nil
}

func assembleCommentAndUser(result *[]dto.Comment, myId int64, authorId int64, commentDTO dto.CommentDTO) {
	user := GetUserInfo(myId, authorId)
	comment := dto.Comment{
		User:       user,
		CommentDTO: commentDTO,
	}
	*result = append(*result, comment)
}
