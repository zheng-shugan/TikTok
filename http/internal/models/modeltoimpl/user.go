package modeltoimpl

import (
	"errors"
	"github.com/sunflower10086/TikTok/http/internal/models"
	"github.com/sunflower10086/TikTok/http/internal/video"
)

func MapUser(modelUser *models.User) (*video.User, error) {
	if modelUser == nil {
		return nil, errors.New("modelUser-to-implUser: 数据为空")
	}

	return &video.User{
		ID:              modelUser.ID,
		Name:            modelUser.Username,
		FollowCount:     modelUser.OtherInfo.FollowCount,
		FollowerCount:   modelUser.OtherInfo.FollowerCount,
		IsFollow:        false, // 默认没有被关注
		Avatar:          modelUser.Avatar,
		BackgroundImage: modelUser.BackgroundImage,
		Signature:       modelUser.Signature,
		TotalFavorited:  modelUser.OtherInfo.TotalFavorited,
		WorkCount:       modelUser.OtherInfo.WorkCount,
		FavoriteCount:   modelUser.OtherInfo.FavoriteCount,
	}, nil
}
