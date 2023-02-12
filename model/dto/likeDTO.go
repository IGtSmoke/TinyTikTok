package dto

type LikeDTO struct {
	UserId  int64
	VideoId int64
	IsThumb int8 `gorm:"column:status"`
}
