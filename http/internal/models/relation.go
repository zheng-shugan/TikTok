package models

type Follow struct {
	Id         int64 `json:"id"`
	UserId     int64 `json:"user_id"`
	FollowerId int64 `json:"follower_id"`
	Type       int8  `json:"type"`
}
