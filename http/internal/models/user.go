package models

import (
	"gorm.io/gorm"
)

// User 用户的model，与数据库保持一致
//type User struct {
//	ID              int64  `json:"id"`
//	UserName        string `gorm:"column:username" json:"name"` // 用户名称
//	Password        string `json:"password,omitempty"`
//	Avatar          string `json:"avatar"`           // 用户头像
//	BackgroundImage string `json:"background_image"` // 用户个人页顶部大图
//	FollowCount     int64  `json:"follow_count"`     // 关注总数
//	FollowerCount   int64  `json:"follower_count"`   // 粉丝总数
//	IsFollow        bool   `json:"is_follow"`        // true-已关注，false-未关注
//	FavoriteCount   int64  `json:"favorite_count"`   // 喜欢数
//	Signature       string `json:"signature"`        // 个人简介
//	TotalFavorited  string `json:"total_favorited"`  // 获赞数量
//	WorkCount       int64  `json:"work_count"`       // 作品数
//}

// TableName 实现Gorm的接口
//func (*User) TableName() string {
//	return constants.UserTableName
//}

type User struct {
	gorm.Model
	Username        string `gorm:"type:varchar(255);not null;unique" json:"username" binding:"required"`
	Password        string `gorm:"type:varchar(255);not null" json:"password" binding:"required"`
	Avatar          string `gorm:"type:varchar(255)" json:"avatar"`
	BackgroundImage string `gorm:"type:varchar(255)" json:"background_image"`
	Signature       string `gorm:"type:varchar(255)" json:"signature"`
}
