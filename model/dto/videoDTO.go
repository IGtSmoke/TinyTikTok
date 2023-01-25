package dto

// VideoDTO video基本信息
type VideoDTO struct {
	ID       uint `gorm:"primarykey"` //videoId由数据库生成
	AuthorID int64
	Title    string `json:"title"`                       // 视频标题
	CoverURL string `column:"coverUrl" json:"cover_url"` // 视频封面地址
	PlayURL  string `column:"playUrl" json:"play_url"`
}
