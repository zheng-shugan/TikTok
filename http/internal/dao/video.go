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
	var user *video.User               // 业务层

	conn := db.GetDB().WithContext(ctx)

	// 查询发布列表
	err := conn.Find(&videos, "author_id = ?", userID).Error
	if err != nil {
		return nil, err
	}
	// 查询发布者信息
	err = conn.Preload("OtherInfo").Find(&user, "id = ?", userID).Error
	if err != nil {
		return nil, err
	}

	// 将数据层映射到业务层
	for _, v := range videos {
		var v3 *video.Video // 业务层
		v3, err = modeltoimpl.MapVideo(v)
		if err != nil {
			return nil, err
		}
		v3.Author = user
		videos2 = append(videos2, v3)
	}

	return videos2, nil
}

func CalFavoriteCount(ctx context.Context, videoID int64) (int64, error) {
	conn := db.GetDB().WithContext(ctx)
	var count int64
	err := conn.Table("user_favorite").Where("video_id = ?", videoID).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

func CalCommentCount(ctx context.Context, videoID int64) (int64, error) {
	conn := db.GetDB().WithContext(ctx)
	var count int64
	err := conn.Model(&models.Comment{}).Where("video_id = ?", videoID).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

func QueryFeedVideo(ctx context.Context, limit int, latestTime int64) ([]*video.Video, error) {
	videos := make([]*models.Video, limit) // 数据层
	videos2 := make([]*video.Video, 0)     // 业务层

	conn := db.GetDB().WithContext(ctx)

	// 返回按投稿时间倒序的视频列表，视频数由服务端控制，单次最多30个
	err := conn.Preload("User.OtherInfo").
		Preload("User").
		Limit(limit).Order("created_at desc").
		Find(&videos, "created_at <= ?", time.Unix(latestTime, 0)).Error

	if err != nil {
		return nil, err
	}

	// 统计每个视频的点赞数和评论数并将数据层映射到业务层
	for _, v := range videos {
		var user2 *video.User // 业务层
		var v3 *video.Video   // 业务层

		if v == nil {
			continue
		}

		favoriteCount, err := CalFavoriteCount(ctx, v.ID) // 统计视频点赞数
		if err != nil {
			return nil, err
		}

		commentCount, err := CalCommentCount(ctx, v.ID) // 统计视频评论数
		if err != nil {
			return nil, err
		}

		v.FavoriteCount = favoriteCount
		v.CommentCount = commentCount

		v3, err = modeltoimpl.MapVideo(v)
		if err != nil {
			return nil, err
		}
		user2, err = modeltoimpl.MapUser(&v.User)
		if err != nil {
			return nil, err
		}

		v3.Author = user2
		videos2 = append(videos2, v3)
	}

	return videos2, nil
}

func CheckIsFavorite(ctx context.Context, videos []*video.Video, userID int64) error {
	conn := db.GetDB().WithContext(ctx)

	//判断当前视频是否被当前用户点赞
	for _, v := range videos {
		var count int64 = 0
		err := conn.Table("user_favorite").Where("user_id = ? and video_id = ?", userID, v.ID).Count(&count).Error
		if err != nil {
			return err
		}
		if count != 0 {
			v.IsFavorite = true
		} else {
			v.IsFavorite = false
		}
	}

	return nil
}

func CheckIsFollow(ctx context.Context, videos []*video.Video, userID int64) error {
	conn := db.GetDB().WithContext(ctx)

	// 判断视频作者是否被当前用户关注
	for _, v := range videos {
		var count int64 = 0
		err := conn.Table("user_follower").Where("user_id = ? and follower_id = ?", v.Author.ID, userID).Count(&count).Error
		if err != nil {
			return err
		}
		if count != 0 {
			v.Author.IsFollow = true
		} else {
			v.Author.IsFollow = false
		}
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
