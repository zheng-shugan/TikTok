package models

import "github.com/sunflower10086/TikTok/http/pkg/constants"

type UserFavorite struct {
	UserID  int64 `json:"user_id"`
	VideoID int64 `json:"video_id"`
}

func (*UserFavorite) TableName() string {
	return constants.UserFavoriteTableName
}
