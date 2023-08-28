package modelToimpl

import (
	"github.com/sunflower10086/TikTok/interaction/internal/dao/models"
	interaction "github.com/sunflower10086/TikTok/interaction/pb"
	"strconv"
)

func MapFavorite(v *models.Video) *interaction.Video {
	return &interaction.Video{
		Author: &interaction.User{
			ID:              v.User.ID,
			Name:            v.User.Username,
			FollowCount:     v.User.OtherInfo.FollowCount,
			FollowerCount:   v.User.OtherInfo.FollowerCount,
			Avatar:          v.User.Avatar,
			BackgroundImage: v.User.BackgroundImage,
			Signature:       v.User.Signature,
			TotalFavorited:  strconv.FormatInt(v.User.OtherInfo.TotalFavorited, 10),
			WorkCount:       v.User.OtherInfo.WorkCount,
			FavoriteCount:   v.User.OtherInfo.FavoriteCount,
		},
		PlayURL:       v.PlayURL,
		CoverURL:      v.CoverURL,
		Title:         v.Title,
		ID:            v.ID,
		CommentCount:  v.CommentCount,
		FavoriteCount: v.FavoriteCount,
	}
}
