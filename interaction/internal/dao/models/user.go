package models

type User struct {
	CustomBaseModel
	Username        string    `gorm:"type:varchar(255);not null;unique" json:"username" binding:"required"`
	Password        string    `gorm:"type:varchar(255);not null" json:"password" binding:"required"`
	Avatar          string    `gorm:"type:varchar(255)" json:"avatar"`
	BackgroundImage string    `gorm:"type:varchar(255)" json:"background_image"`
	Signature       string    `gorm:"type:varchar(255)" json:"signature"`
	OtherInfoID     int64     // 其他信息的id，为外键
	OtherInfo       OtherInfo // belongs to 关系

	Follower []User `gorm:"many2many:user_follower;"` // 自引用的many to many关系，表示user的跟随着

	FavoriteVideo []Video `gorm:"many2many:user_favorite"`
}

type OtherInfo struct {
	ID             int64 `gorm:"primaryKey" json:"id"`
	FollowCount    int64 `gorm:"type:bigint;default:0" json:"follow_count"`
	FollowerCount  int64 `gorm:"type:bigint;default:0" json:"follower_count"`
	FavoriteCount  int64 `gorm:"type:bigint;default:0" json:"favorite_count"`
	TotalFavorited int64 `gorm:"type:bigint;default:0" json:"total_favorited"`
	WorkCount      int64 `gorm:"type:bigint;default:0" json:"work_count"`
}
