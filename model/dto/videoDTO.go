package dto

// VideoDTO video基本信息
type VideoDTO struct {
	Id       uint   `gorm:"primarykey" json:"id"` // videoId由数据库生成
	AuthorID int64  `json:"-"`
	Title    string `json:"title"`                        // 视频标题
	CoverURL string `column:"cover_url" json:"cover_url"` // 视频封面地址
	PlayURL  string `column:"play_url" json:"play_url"`
}
