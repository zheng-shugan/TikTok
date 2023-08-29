package rpcTohttp

//import (
//	"github.com/sunflower10086/TikTok/http/internal/interaction"
//	"strconv"
//)

//func MapHttp(v *___interaction.Video) *interaction.Video {
//	return &interaction.Video{
//		Author: &interaction.User{
//			ID:              v.Author.ID,
//			Name:            v.Author.Username,
//			FollowCount:     v.Author.OtherInfo.FollowCount,
//			FollowerCount:   v.Author.OtherInfo.FollowerCount,
//			Avatar:          v.Author.Avatar,
//			BackgroundImage: v.Author.BackgroundImage,
//			Signature:       v.Author.Signature,
//			TotalFavorited:  strconv.FormatInt(v.Author.OtherInfo.TotalFavorited, 10),
//			WorkCount:       v.Author.OtherInfo.WorkCount,
//			FavoriteCount:   v.Author.OtherInfo.FavoriteCount,
//		},
//		PlayURL:       v.PlayURL,
//		CoverURL:      v.CoverURL,
//		Title:         v.Title,
//		ID:            v.ID,
//		CommentCount:  v.CommentCount,
//		FavoriteCount: v.FavoriteCount,
//	}
//}
