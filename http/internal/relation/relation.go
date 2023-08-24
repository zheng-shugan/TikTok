package relation

import (
	"github.com/sunflower10086/TikTok/http/internal/models"
	"github.com/sunflower10086/TikTok/http/internal/pkg/result"
)

type Service interface {
}

type ActionRequest struct {
	ToUserID   int64 `json:"to_user_id" binding:"required"`
	ActionType int   `json:"action_type" binding:"required"`
}

type ActionResponse struct {
	result.Response
}

type FollowListRequest struct {
	UserID int64 `json:"user_id"`
}
type FollowerListRequest = FollowListRequest
type FriendListRequest = FollowListRequest

type FollowListResponse struct {
	result.Response
	UserList []models.User `json:"user_list"`
}
type FollowerListResponse = FollowListResponse
type FriendListResponse = FollowListResponse
