package models

import "github.com/sunflower10086/TikTok/http/pkg/constants"

type UserFollow struct {
	UserID     int64 `json:"user_id"`
	FollowerID int64 `json:"follower_id"`
}

func (*UserFollow) TableName() string {
	return constants.UserFollowTableName
}
