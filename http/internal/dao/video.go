package dao

import (
	"context"
	"time"

	"github.com/sunflower10086/TikTok/http/internal/dao/db"
	"github.com/sunflower10086/TikTok/http/internal/models"
	"github.com/sunflower10086/TikTok/http/internal/video"
)

func QueryPublishList(ctx context.Context, userID int64) ([]*video.Video, error) {
	return nil, nil
}

func QueryFeedVideo(ctx context.Context, limit int, latestTime int64) ([]*video.Video, error) {
	return nil, nil
}

func CheckIsFavorite(ctx context.Context, videos []*video.Video, userID int64) error {
	return nil
}

func CheckIsFollow(ctx context.Context, videos []*video.Video, userID int64) error {
	return nil
}

func SaveVideo(ctx context.Context, downUrl, title string, userId int64) error {
	var video models.Video
	video.AuthorID = userId
	video.PlayURL = downUrl
	video.Title = title
	video.CreateTime = time.Now().Unix()
	return db.GetDB().Create(&video).Error
}
