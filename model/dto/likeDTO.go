package dto

type LikeDTO struct {
	UserId  int64
	VideoId int64
	IsThumb bool `column:"cancel"`
}
