package modelToimpl

import (
	"github.com/sunflower10086/TikTok/interaction/internal/dao/models"
	interaction "github.com/sunflower10086/TikTok/interaction/pb"
	"strconv"
)

func MapFavorite(v []*models.Video, v1 []*interaction.Video) {
	for i, _ := range v {
		v1[i].ID = v[i].ID
		v1[i].Author.ID = v[i].User.ID
		v1[i].Author.Name = v[i].User.Username
		v1[i].Author.FollowCount = v[i].User.OtherInfo.FollowCount
		v1[i].Author.FollowerCount = v[i].User.OtherInfo.FollowerCount
		v1[i].Author.Avatar = v[i].User.Avatar
		v1[i].Author.BackgroundImage = v[i].User.BackgroundImage
		v1[i].Author.Signature = v[i].User.Signature
		v1[i].Author.TotalFavorited = strconv.FormatInt(v[i].User.OtherInfo.TotalFavorited, 10)
		v1[i].Author.WorkCount = v[i].User.OtherInfo.WorkCount
		v1[i].Author.FavoriteCount = v[i].User.OtherInfo.FavoriteCount
		v1[i].PlayURL = v[i].PlayURL
		v1[i].CoverURL = v[i].CoverURL
		v1[i].Title = v[i].Title
	}
}
