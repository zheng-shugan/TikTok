package modeltoimpl

import (
	"errors"
	"github.com/sunflower10086/TikTok/http/internal/models"
	"github.com/sunflower10086/TikTok/http/internal/video"
	"regexp"
	"strconv"
)

func strToint(s *string) (int64, error) {
	// 定义正则表达式匹配浮点数和汉字
	re := regexp.MustCompile(`([\d.]+)([万亿])`)
	//
	// 使用正则表达式匹配
	matches := re.FindStringSubmatch(*s)
	if len(matches) != 3 {
		return 0, errors.New("modelUser-to-implUser:正则匹配错误")
	}
	// 解析浮点数
	number, err := strconv.ParseFloat(matches[1], 64)
	if err != nil {
		return 0, errors.New("modelUser-to-implUser: 解析浮点数错误")
	}

	if matches[2] == "万" {
		number = number * 1e4
	} else if matches[2] == "亿" {
		number = number * 1e9
	}

	return int64(number), nil
}

// FIXME: 修改实现
func MapUser(modelUser *models.User) (*video.User, error) {
	//if modelUser == nil {
	//	return nil, errors.New("modelUser-to-implUser: 数据为空")
	//}
	//
	//totalFavorited, err := strToint(&modelUser.TotalFavorited)
	//if err != nil {
	//	return nil, err
	//}
	//
	//return &video.User{
	//	ID:              modelUser.ID,
	//	Name:            modelUser.UserName,
	//	FollowCount:     modelUser.FollowCount,
	//	FollowerCount:   modelUser.FollowerCount,
	//	IsFollow:        modelUser.IsFollow,
	//	Avatar:          modelUser.Avatar,
	//	BackgroundImage: modelUser.BackgroundImage,
	//	Signature:       modelUser.Signature,
	//	TotalFavorited:  totalFavorited,
	//	WorkCount:       modelUser.WorkCount,
	//	FavoriteCount:   modelUser.FavoriteCount,
	//}, nil

	return nil, nil
}
