package impl

import (
	"TinyTikTok/conf"
	"TinyTikTok/dao"
	"TinyTikTok/model/dto"
	"TinyTikTok/utils"
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"mime/multipart"
	"strings"
)

type PublishServiceImpl struct {
}

// Action 上传视频
func (p PublishServiceImpl) Action(c *gin.Context) {
	//获取相关信息
	value, exists := c.Get("userId")
	if exists == false {
		log.Error().Msg("[PublishServiceImpl]Action方法中userDTO对象不存在")
	}
	userId := value.(int64)
	title := c.PostForm("title")
	data, err := c.FormFile("data")
	if err != nil {
		log.Err(err)
		utils.Fail(c, "上传文件失败")
		return
	}
	file, err := data.Open()
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

	err = utils.UploadFile(conf.Conf.VideoBucketName, videoName, file, data.Size)
	if err != nil {
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
	err = utils.UploadFile(conf.Conf.CoverBucketName, coverName, coverReader, int64(len(coverData)))

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
	//保存视频信息到mysql
	dao.SaveVideo(videoDTO)
	response := dto.Result{}
	utils.InitSuccessResult(&response)
	utils.Success(c, response)
	return
}
