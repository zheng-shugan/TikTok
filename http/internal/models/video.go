package models

// Video 视频的model,与数据库保持一致
type Video struct {
	CustomBaseModel
	PlayURL       string `gorm:"type:varchar(255)" json:"play_url"`
	CoverURL      string `gorm:"type:varchar(255)" json:"cover_url"`
	FavoriteCount int64  `gorm:"type:bigint" json:"favorite_count"`
	CommentCount  int64  `gorm:"type:bigint" json:"comment_count"`
	Title         string `gorm:"type:varchar(255)" json:"title"`
	AuthorID      int64  `json:"author_id"`
	User          User   `gorm:"foreignKey:AuthorID"`
}
