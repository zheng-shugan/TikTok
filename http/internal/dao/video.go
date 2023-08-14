package dao

import (
	"context"

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

func SaveVideo(ctx context.Context, video video.Video) error {
	return nil
}
