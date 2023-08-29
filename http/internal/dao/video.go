package dao

import (
	"context"
	"fmt"
	"time"

	"github.com/sunflower10086/TikTok/http/internal/dao/db"
	"github.com/sunflower10086/TikTok/http/internal/models"
	"github.com/sunflower10086/TikTok/http/internal/models/modeltoimpl"
	"github.com/sunflower10086/TikTok/http/internal/video"
)

// 查询发布列表
func QueryPublishList(ctx context.Context, userID int64) ([]*video.Video, error) {
	videos := make([]*models.Video, 0) // 数据层
	videos2 := make([]*video.Video, 0) // 业务层
	var user *models.User              // 业务层

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
	// 获取实时User信息
	err = RealTimeUser(ctx, user)
	if err != nil {
		return nil, err
	}

	// 获取每个视频的实时信息并将数据层映射到业务层
	for _, v := range videos {
		var v3 *video.Video // 业务层

		// 获取实时Video信息
		err = RealTimeVideo(ctx, v)
		if err != nil {
			return nil, err
		}

		// 映射到业务层
		v3, err = modeltoimpl.MapVideo(v)
		if err != nil {
			return nil, err
		}
		v3.Author, err = modeltoimpl.MapUser(user)
		if err != nil {
			return nil, err
		}
		videos2 = append(videos2, v3)
	}

	return videos2, nil
}

// 获取视频流
func QueryFeedVideo(ctx context.Context, limit int, latestTime int64) ([]*video.Video, error) {
	videos := make([]*models.Video, limit) // 数据层
	videos2 := make([]*video.Video, 0)     // 业务层

	conn := db.GetDB().WithContext(ctx)

	// 返回按投稿时间倒序的视频列表，视频数由服务端控制，单次最多30个
	err := conn.Preload("User.OtherInfo").
		Preload("User").
		Limit(limit).Order("created_at desc").
		Where("created_at <= ?", latestTime).
		Find(&videos).Error

	if err != nil {
		return nil, err
	}

	// 获取每个视频的实时信息并将数据层映射到业务层
	for _, v := range videos {
		var user2 *video.User // 业务层
		var v3 *video.Video   // 业务层

		if v == nil {
			continue
		}

		// 获取实时的video信息和User信息
		err = RealTimeVideo(ctx, v)
		if err != nil {
			return nil, err
		}
		err = RealTimeUser(ctx, &v.User)
		if err != nil {
			return nil, err
		}

		// 映射到业务层
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

// 保存用户发布的视频
func SaveVideo(ctx context.Context, downUrl, title string, userId int64) error {
	var v models.Video
	fmt.Println(userId)
	v.User = models.User{
		CustomBaseModel: models.CustomBaseModel{ID: userId},
	}
	v.AuthorID = userId
	v.PlayURL = downUrl
	v.Title = title
	v.CreatedAt = time.Now().Unix()
	return db.GetDB().Create(&v).Error
}
