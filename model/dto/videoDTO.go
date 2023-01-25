package dto

// VideoDTO video基本信息
type VideoDTO struct {
	ID       uint `gorm:"primarykey"` // videoId由数据库生成
	AuthorID int64
	Title    string `json:"title"`                       // 视频标题
	CoverURL string `column:"cover_url" json:"coverUrl"` // 视频封面地址
	PlayURL  string `column:"play_url" json:"playUrl"`
}
