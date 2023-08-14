package dao

import (
	"context"
	"github.com/sunflower10086/TikTok/http/internal/videostream"
)

func QueryFeedVideo(ctx context.Context, limit int, latestTime int64) ([]*videostream.Video, error) {
	return nil, nil
}

func CheckIsFavorite(ctx context.Context, videos []*videostream.Video, userID int64) error {
	return nil
}

func CheckIsFollow(ctx context.Context, videos []*videostream.Video, userID int64) error {
	return nil
}
