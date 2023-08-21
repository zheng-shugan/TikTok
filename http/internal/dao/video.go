package dao

import (
	"context"
	"time"

	"github.com/sunflower10086/TikTok/http/internal/dao/db"
	"github.com/sunflower10086/TikTok/http/internal/models"
	"github.com/sunflower10086/TikTok/http/internal/models/modeltoimpl"
	"github.com/sunflower10086/TikTok/http/internal/video"
)

func QueryPublishList(ctx context.Context, userID int64) ([]*video.Video, error) {
	videos := make([]*models.Video, 0) // 数据层
	videos2 := make([]*video.Video, 0) // 业务层
	conn := db.GetDB().WithContext(ctx)

	// 查询发布列表
	err := conn.Find(&videos, "author_id = ?", userID).Error
	if err != nil {
		return nil, err
	}
	for i, v := range videos {
		var err2 error
		videos2[i], err2 = modeltoimpl.MapVideo(v)
		if err2 != nil {
			return nil, err2
		}
	}

	return videos2, nil
}

func QueryFeedVideo(ctx context.Context, limit int, latestTime int64) ([]*video.Video, error) {
	videos := make([]*models.Video, limit) // 数据层model
	videos2 := make([]*video.Video, limit) // 业务层数据
	conn := db.GetDB().WithContext(ctx)

	// 返回按投稿时间倒序的视频列表，视频数由服务端控制，单次最多30个
	err := conn.Limit(limit).Order("create_time desc").Find(&videos, "create_time <= ?", latestTime).Error
	if err != nil {
		return nil, err
	}
	// 查询视频作者信息并打包数据
	for i, v := range videos {
		var user models.User  // 数据层
		var user2 *video.User // 业务层
		err := conn.Find(&user, "id = ?", v.AuthorID).Error
		if err != nil {
			return nil, err
		}
		videos2[i], err = modeltoimpl.MapVideo(v)
		if err != nil {
			return nil, err
		}
		user2, err = modeltoimpl.MapUser(&user)
		if err != nil {
			return nil, err
		}
		videos2[i].Author = user2
	}

	return videos2, nil
}

func CheckIsFavorite(ctx context.Context, videos []*video.Video, userID int64) error {
	var userFavorite models.UserFavorite
	userFavorite.VideoID = -1
	conn := db.GetDB().WithContext(ctx)

	for _, v := range videos {
		err := conn.Find(&userFavorite, "user_id = ? and video_id = ?", userID, v.ID).Error
		if err != nil {
			return err
		}
		if userFavorite.UserID != -1 {
			v.IsFavorite = true
		} else {
			v.IsFavorite = false
		}
		userFavorite.VideoID = -1
	}

	return nil
}

func CheckIsFollow(ctx context.Context, videos []*video.Video, userID int64) error {
	var userFollow models.UserFollow
	userFollow.FollowerID = -1
	conn := db.GetDB().WithContext(ctx)

	for _, v := range videos {
		err := conn.Find(&userFollow, "user_id = ? and follower_id = ?", userID, v.Author.ID).Error // TODO:命名待确定
		if err != nil {
			return err
		}
		if userFollow.FollowerID != -1 {
			v.Author.IsFollow = true
		} else {
			v.Author.IsFollow = false
		}
		userFollow.FollowerID = -1
	}

	return nil
}

func SaveVideo(ctx context.Context, downUrl, title string, userId int64) error {
	var video models.Video
	video.AuthorID = userId
	video.PlayURL = downUrl
	video.Title = title
	video.CreatedAt = time.Now()
	return db.GetDB().Create(&video).Error
}
