package modeltoimpl

import (
	"errors"
	"github.com/sunflower10086/TikTok/http/internal/models"
	"github.com/sunflower10086/TikTok/http/internal/video"
)

func MapVideo(modelVideo *models.Video) (*video.Video, error) {
	if modelVideo == nil {
		return nil, errors.New("modelVideo-to-implVideo: 数据为空")
	}

	return &video.Video{
		ID:            modelVideo.ID,
		PlayUrl:       modelVideo.PlayURL,
		CoverUrl:      modelVideo.CoverURL,
		FavoriteCount: modelVideo.FavoriteCount,
		CommentCount:  modelVideo.CommentCount,
		IsFavorite:    modelVideo.IsFavorite,
		Title:         modelVideo.Title,
		PublishTime:   modelVideo.CreateTime,
	}, nil
}
