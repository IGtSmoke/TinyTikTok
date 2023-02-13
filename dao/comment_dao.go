package dao

import (
	"TinyTikTok/conf/setup"
	"TinyTikTok/model/dto"
	"gorm.io/gorm"
)

type CommentPO struct {
	gorm.Model
	dto.CommentDTO
}

func (c CommentPO) TableName() string {
	return "comments"
}

func QueryCommentById(commentId int64) dto.CommentDTO {
	var Result dto.CommentDTO
	setup.Mdb.Model(CommentPO{}).Where("id = ?", commentId).Find(&Result)
	return Result
}

func DeleteCommentById(commentId int64) {
	setup.Mdb.Delete(CommentPO{}, commentId)
}

func CreateComment(commentDTO dto.CommentDTO) uint {
	var Result dto.CommentDTO
	setup.Mdb.Create(&CommentPO{
		CommentDTO: commentDTO,
	})
	setup.Mdb.Model(CommentPO{}).Where("user_id = ? AND video_id = ? AND comment_text = ?", commentDTO.UserID, commentDTO.VideoID, commentDTO.CommentText).Find(&Result)
	return Result.Id
}

func QueryVideoComment(videoId int64) []dto.CommentDTO {
	var Result []dto.CommentDTO
	setup.Mdb.Model(CommentPO{}).Where("video_id = ?", videoId).Find(&Result)
	return Result
}
