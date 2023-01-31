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
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v9"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"mime/multipart"
	"strconv"
	"strings"
	"time"
)

type PublishServiceImpl struct {
}

// Action 上传视频
func (p PublishServiceImpl) Action(c *gin.Context) {
	//获取相关信息
	userId, err := utils.GetUserIdByMiddleware(c)
	if err != nil {
		log.Err(err).Msg("[PublishServiceImpl]action中获取userId失败")
		utils.Fail(c, "上传文件失败")
		return
	}
	title, err := getTitle(c)
	if err != nil {
		log.Error().Msg(err.Error())
		utils.Fail(c, "上传文件失败")
		return
	}
	data, err := getFile(c)
	if err != nil {
		log.Err(err)
		utils.Fail(c, "上传文件失败")
		return
	}
	file, err := openFile(data)
	if err != nil {
		log.Err(err)
		utils.Fail(c, "上传文件失败")
		return
	}
	//关闭文件
	defer func(file multipart.File) {
		err := file.Close()
		if err != nil {
			log.Err(err)
		}
	}(file)

	//上传视频到minio
	uuId := uuid.NewString()
	videoName := uuId + "." + "mp4"

	if err = utils.UploadFile(conf.Conf.VideoBucketName, videoName, file, data.Size); err != nil {
		log.Err(err)
	}

	//获取视频链接
	url, err := utils.GetFileUrl(conf.Conf.VideoBucketName, videoName, 0)
	if err != nil {
		log.Err(err)
	}
	playUrl := strings.Split(url.String(), "?")[0]

	//上传封面
	coverName := uuId + "." + "jpg"
	coverData, err := utils.ReadFrameAsJpeg(playUrl)
	if err != nil {
		log.Err(err)
	}
	coverReader := bytes.NewReader(coverData)
	if err = utils.UploadFile(conf.Conf.CoverBucketName, coverName, coverReader, int64(len(coverData))); err != nil {
		log.Err(err)
	}

	// 获取封面链接
	tmpCoverUrl, err := utils.GetFileUrl(conf.Conf.CoverBucketName, coverName, 0)
	if err != nil {
		log.Err(err)
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
		log.Err(err)
	}
	//保存视频信息到redis
	tokenKey := utils.PersonalVideosKey + strconv.FormatInt(userId, 10)
	if _, err := setup.Rdb.Pipelined(setup.Rctx, func(rdb redis.Pipeliner) error {
		rdb.LPush(setup.Rctx, tokenKey, videoDTOJson)
		rdb.LPush(setup.Rctx, utils.TotalVideosKey, videoDTOJson)
		return nil
	}); err != nil {
		log.Err(err)
	}
	//保存视频信息到mysql
	dao.SaveVideo(videoDTO)
	response := dto.Result{}
	utils.InitSuccessResult(&response)
	utils.Success(c, response)
	return
}

func (p PublishServiceImpl) List(c *gin.Context) {
	myId, err := utils.GetUserIdByMiddleware(c)
	if err != nil {
		log.Err(err).Msg("[PublishServiceImpl]List中获取userId失败")
		utils.Fail(c, "获取失败")
		return
	}

	userIdStr := c.Query("user_id")
	authorId, err := strconv.ParseInt(userIdStr, 10, 64)

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
				log.Err(err)
			}
			assembleUser(&result, myId, videoDTO)
		}
		listSuccess(c, &result)
		return
	}

	//redis查询失败尝试mysql查询
	log.Err(err)
	videoDTOS := dao.GetVideosByAuthorId(authorId)
	for _, tmp := range videoDTOS {
		assembleUser(&result, myId, tmp)
	}
	listSuccess(c, &result)
	return
}

func listSuccess(c *gin.Context, result *[]dto.Video) {
	response := dto.PublishListResponse{
		VideoList: *result,
	}
	utils.InitSuccessResult(&response.Result)
	utils.Success(c, response)
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
	videoId := video.VideoDTO.ID
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
		log.Err(err)
		//todo 从数据库中获取点赞数
	}
	//获得结果
	commentCount := exec[0].(*redis.IntCmd).Val()
	favoriteCount := exec[1].(*redis.IntCmd).Val()
	isFavorite := exec[2].(*redis.BoolCmd).Val()

	video.CommentCount = commentCount
	video.FavoriteCount = favoriteCount
	video.IsFavorite = isFavorite
}

func (p PublishServiceImpl) Feed(c *gin.Context) {
	lastTime, err := getLastTime(c)
	if err != nil {
		log.Err(err)
		lastTime = time.Now()
	}
	myId, _ := utils.GetUserIdByMiddleware(c)
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
	utils.Success(c, response)
	return
}

func getLastTime(c *gin.Context) (lastTime time.Time, err error) {
	value := c.Query("latest_time")
	if value == "0" {
		lastTime = time.Now()
		return
	}
	inputTime, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		log.Err(err)
		return
	}
	lastTime = time.Unix(inputTime, 0)
	return
}

func getTitle(c *gin.Context) (string, error) {
	title := c.PostForm("title")
	if title == "" {
		return "", errors.New("title不能为空")
	}
	return title, nil
}

func getFile(c *gin.Context) (*multipart.FileHeader, error) {
	data, err := c.FormFile("data")
	if err != nil {
		return nil, err
	}
	return data, nil
}

func openFile(file *multipart.FileHeader) (multipart.File, error) {
	f, err := file.Open()
	if err != nil {
		return nil, err
	}
	return f, nil
}
