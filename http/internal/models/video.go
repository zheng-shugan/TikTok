package models

import "github.com/sunflower10086/TikTok/http/pkg/constants"

// Video 视频的model,与数据库保持一致
type Video struct {
	ID            int64  `json:"id"`
	CreateTime    int64  `json:"create_time"`
	AuthorID      int64  `json:"author_id"`
	PlayURL       string `json:"play_url"`
	CoverURL      string `json:"cover_url"`
	FavoriteCount int64  `json:"favorite_count"`
	CommentCount  int64  `json:"comment_count"`
	IsFavorite    bool   `json:"is_favorite"`
	Title         string `json:"title"`
}

// TableName 实现Gorm的接口
func (*Video) TableName() string {
	return constants.VideoTableName
}
