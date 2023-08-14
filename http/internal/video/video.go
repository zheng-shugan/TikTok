package video

import (
	"context"
	"mime/multipart"

	"github.com/gin-gonic/gin"
	"github.com/sunflower10086/TikTok/http/internal/pkg/result"
)

type Server interface {
	// GetFeedVideo 获得30个视频
	GetFeedVideo(context.Context, *GetFeedVideoReq) (*GetFeedVideoResp, error)
	// PublishAction 上传视频
	PublishAction(*gin.Context, *PublishRequest) (*PublishResponse, error)
	// GetPublishList 获得用户所有投稿过的视频
	GetPublishList(context.Context, GetPublishListReq) (GetPublishListResp, error)
}

type GetPublishListReq struct {
	UserID int64  `json:"user_id"`
	Token  string `json:"token"`
}

type GetPublishListResp struct {
	StatusCode int32    `json:"status_code"`
	StatusMsg  string   `json:"status_msg"`
	VideoList  []*Video `json:"video_list"`
}

type GetFeedVideoReq struct {
	LatestTime int64  `json:"latest_time"`
	Token      string `json:"token"`
}

type GetFeedVideoResp struct {
	StatusCode int32    `json:"status_code" binding:"required"`
	StatusMsg  string   `json:"status_msg"`
	VideoList  []*Video `json:"video_list" binding:"required"`
	NextTime   *int64   `json:"next_time"`
}

type Video struct {
	ID            int64  `json:"id"`
	Author        *User  `json:"author"`
	PlayUrl       string `json:"play_url"`
	CoverUrl      string `json:"cover_url"`
	FavoriteCount int64  `json:"favorite_count"`
	CommentCount  int64  `json:"comment_count"`
	IsFavorite    bool   `json:"is_favorite"`
	Title         string `json:"title"`
	PublishTime   int64  // TODO： 命名待确定
}

type User struct {
	ID              int64  `json:"id"`
	Name            string `json:"name"`
	FollowCount     int64  `json:"follow_count"`
	FollowerCount   int64  `json:"follower_count"`
	IsFollow        bool   `json:"is_follow"`
	Avatar          string `json:"avatar"`
	BackgroundImage string `json:"background_image"`
	Signature       string `json:"signature"`
	TotalFavorited  int64  `json:"total_favorited"`
	WorkCount       int64  `json:"work_count"`
	FavoriteCount   int64  `json:"favorite_count"`
}

type PublishRequest struct {
	Data  *multipart.FileHeader `json:"data" binding:"required" form:"data"`
	Title string                `json:"title" binding:"required" form:"title"`
}

type PublishResponse struct {
	result.Response
}
