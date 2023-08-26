package dao

import (
	"context"
)

func AddFavorite(ctx context.Context, userID int64, videoID int64) error {
	return nil
}

func DelFavorite(ctx context.Context, userID int64, videoID int64) error {
	return nil
}

//func GetFavoriteList(ctx context.Context, userID int64) ([]*interaction.Video, error) {
//	return nil, nil
//}
