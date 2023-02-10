package impl

import (
	"TinyTikTok/conf"
	"TinyTikTok/conf/setup"
	"TinyTikTok/dao"
	"TinyTikTok/model/dto"
	"TinyTikTok/utils"
	"bytes"
	"encoding/json"
	"errors"
	"github.com/go-redis/redis/v9"
	"github.com/google/uuid"
	"mime/multipart"
	"runtime/debug"
	"strconv"
	"strings"
)

type PublishServiceImpl struct {
}

// Action 上传视频
func (p PublishServiceImpl) Action(userId int64, title string, data *multipart.FileHeader) (dto.Result, error) {
	response := dto.Result{}
	file, err := openFile(data)
	if err != nil {
		return response, errors.New("文件打开失败")
	}
	//关闭文件
	defer func(file multipart.File) {
		err := file.Close()
		if err != nil {
			setup.Logger("common").Err(err).Send()
		}
	}(file)
	//上传视频到minio
	uuId := uuid.NewString()
	videoName := uuId + "." + "mp4"

	if err = utils.UploadFile(conf.Conf.VideoBucketName, videoName, file, data.Size); err != nil {
		return response, errors.New("上传文件失败")
	}

	//获取视频链接
	url, err := utils.GetFileUrl(conf.Conf.VideoBucketName, videoName, 0)
	if err != nil {
		return response, errors.New("获取视频连接失败")
	}
	playUrl := strings.Split(url.String(), "?")[0]

	//上传封面
	coverName := uuId + "." + "jpg"
	coverData, err := utils.ReadFrameAsJpeg(playUrl)
	if err != nil {
		return response, errors.New("封面截取失败")
	}
	coverReader := bytes.NewReader(coverData)
	if err = utils.UploadFile(conf.Conf.CoverBucketName, coverName, coverReader, int64(len(coverData))); err != nil {
		return response, errors.New("上传封面失败")
	}

	// 获取封面链接
	tmpCoverUrl, err := utils.GetFileUrl(conf.Conf.CoverBucketName, coverName, 0)
	if err != nil {
		return response, errors.New("获取封面连接失败")
	}
	coverUrl := strings.Split(tmpCoverUrl.String(), "?")[0]

	videoDTO := dto.VideoDTO{
		AuthorID: userId,
		Title:    title,
		CoverURL: coverUrl,
		PlayURL:  playUrl,
	}
	videoDTOJson, err := json.Marshal(videoDTO)
	if err != nil {
		return response, errors.New("videoDTO json化失败")
	}
	//保存视频信息到redis
	tokenKey := utils.PersonalVideosKey + strconv.FormatInt(userId, 10)
	if _, err := setup.Rdb.Pipelined(setup.Rctx, func(rdb redis.Pipeliner) error {
		rdb.LPush(setup.Rctx, tokenKey, videoDTOJson)
		rdb.LPush(setup.Rctx, utils.TotalVideosKey, videoDTOJson)
		return nil
	}); err != nil {
		return response, errors.New("保存到redis失败")
	}
	//保存视频信息到mysql
	dao.SaveVideo(videoDTO)
	utils.InitSuccessResult(&response)
	return response, nil
}

func (p PublishServiceImpl) List(myId int64, userIdStr string) (dto.PublishListResponse, error) {
	response := dto.PublishListResponse{}
	authorId, err := strconv.ParseInt(userIdStr, 10, 64)
	if err != nil {
		setup.Logger("common").Err(err).Interface("stack", string(debug.Stack())).Send()
	}
	//从redis中查询作者视频信息
	tokenKey := utils.PersonalVideosKey + userIdStr
	val, err := setup.Rdb.LRange(setup.Rctx, tokenKey, 0, -1).Result()
	result := make([]dto.Video, 0, len(val))

	// 如果从redis中获取成功，则直接返回结果
	if err != nil {
		var videoDTO dto.VideoDTO
		for _, tmp := range val {
			//获取videoDTO
			if err := json.Unmarshal([]byte(tmp), &videoDTO); err != nil {
				setup.Logger("common").Err(err).Interface("stack", string(debug.Stack())).Send()
			}
			assembleUser(&result, myId, videoDTO)
		}
		response.VideoList = result
		utils.InitSuccessResult(&response.Result)
		return response, nil
	}

	//redis查询失败尝试mysql查询
	setup.Logger("common").Err(err).Send()
	videoDTOS := dao.GetVideosByAuthorId(authorId)
	for _, tmp := range videoDTOS {
		assembleUser(&result, myId, tmp)
	}
	response.VideoList = result
	utils.InitSuccessResult(&response.Result)
	return response, nil
}

func assembleUser(result *[]dto.Video, myId int64, videoDTO dto.VideoDTO) {
	user := GetUserInfo(myId, videoDTO.AuthorID)
	//组装
	video := dto.Video{
		Author:   user,
		VideoDTO: videoDTO,
	}
	//获取评论和点赞数
	getVideoCommentAndFavouriteCountInfo(&video, myId)
	*result = append(*result, video)
}

// 查询视频评论数、点赞数和自身是否点赞
func getVideoCommentAndFavouriteCountInfo(video *dto.Video, myId int64) {
	videoId := video.VideoDTO.Id
	//根据videoId查询redis获得评论数和点赞数
	commentKey := utils.VideoCommentKey + strconv.FormatInt(int64(videoId), 10)
	favoriteKey := utils.VideoLikeKey + strconv.FormatInt(int64(videoId), 10)
	//构建管道
	pipeline := setup.Rdb.Pipeline()
	pipeline.SCard(setup.Rctx, commentKey)
	pipeline.SCard(setup.Rctx, favoriteKey)
	pipeline.SIsMember(setup.Rctx, favoriteKey, myId)
	exec, err := pipeline.Exec(setup.Rctx)
	if err != nil {
		setup.Logger("common").Err(err).Send()
	}
	//获得结果
	commentCount := exec[0].(*redis.IntCmd).Val()
	favoriteCount := exec[1].(*redis.IntCmd).Val()
	isFavorite := exec[2].(*redis.BoolCmd).Val()

	video.CommentCount = commentCount
	video.FavoriteCount = favoriteCount
	video.IsFavorite = isFavorite
}

func openFile(file *multipart.FileHeader) (multipart.File, error) {
	f, err := file.Open()
	if err != nil {
		return nil, err
	}
	return f, nil
}
