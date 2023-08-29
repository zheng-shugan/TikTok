package rpcTohttp

import (
	"github.com/sunflower10086/TikTok/http/internal/interaction"
	___interaction "github.com/sunflower10086/TikTok/interaction/pb"
	"strconv"
)

func MapHttp(v *___interaction.Video) *interaction.Video {
	totalFavorited, _ := strconv.ParseInt(v.Author.TotalFavorited, 10, 64)

	return &interaction.Video{
		Author: &interaction.User{
			ID:              v.Author.ID,
			Name:            v.Author.Name,
			FollowCount:     v.Author.FollowCount,
			FollowerCount:   v.Author.FollowerCount,
			Avatar:          v.Author.Avatar,
			BackgroundImage: v.Author.BackgroundImage,
			Signature:       v.Author.Signature,
			TotalFavorited:  totalFavorited,
			WorkCount:       v.Author.WorkCount,
			FavoriteCount:   v.Author.FavoriteCount,
			IsFollow:        v.Author.IsFollow,
		},
		PlayUrl:       v.PlayURL,
		CoverUrl:      v.CoverURL,
		Title:         v.Title,
		ID:            v.ID,
		CommentCount:  v.CommentCount,
		FavoriteCount: v.FavoriteCount,
		IsFavorite:    v.IsFavorite,
	}
}
