package dao

import (
	"context"
	"github.com/sunflower10086/TikTok/interaction/internal/dao/db"
	"github.com/sunflower10086/TikTok/interaction/internal/dao/models"
	"github.com/sunflower10086/TikTok/interaction/internal/dao/models/modelToimpl"
	interaction "github.com/sunflower10086/TikTok/interaction/pb"
)

type Favorite struct {
	UserID  int64
	VideoID int64
}

// 点赞
func AddFavorite(ctx context.Context, userID int64, videoID int64) error {
	conn := db.GetDB().WithContext(ctx)

	// 检查表中是否已存在该数据
	check, err := CheckIsFavorite(ctx, videoID, userID)
	if err != nil {
		return err
	}

	// 如果不存在则插入新数据
	if !check {
		err = conn.Table("user_favorite").Create(&Favorite{
			UserID:  userID,
			VideoID: videoID,
		}).Error

		if err != nil {
			return err
		}
	}

	return nil
}

// 取消点赞
func DelFavorite(ctx context.Context, userID int64, videoID int64) error {
	conn := db.GetDB().WithContext(ctx)

	// 检查表中是否已存在该数据
	check, err := CheckIsFavorite(ctx, videoID, userID)
	if err != nil {
		return err
	}

	// 如果存在则删除数据
	if check {
		err = conn.Table("user_favorite").
			Where("user_id = ? and video_id = ?", userID, videoID).
			Delete(&Favorite{}).Error
		if err != nil {
			return err
		}
	}

	return nil
}

// 获取点赞列表
func GetFavoriteList(ctx context.Context, userID int64) ([]*interaction.Video, error) {
	var videoID []int64
	videos := make([]*models.Video, 0) // 数据层

	conn := db.GetDB().WithContext(ctx)

	// 查询用户点赞视频的ID
	err := conn.Table("user_favorite").
		Where("user_id = ?", userID).
		Find(&Favorite{}).
		Pluck("video_id", &videoID).Error

	if err != nil {
		return nil, err
	}

	// 根据视频ID查询对应视频
	err = conn.Preload("User.OtherInfo").
		Preload("User").
		Where("video_id in ?", videoID).
		Find(&videos).Error

	videos2 := make([]*interaction.Video, len(videos)) // 业务层

	// 获取实时的Video和User信息
	for _, v := range videos {
		err = RealTimeVideo(ctx, v)
		if err != nil {
			return nil, err
		}
		err = RealTimeUser(ctx, &v.User)
		if err != nil {
			return nil, err
		}
	}

	// 数据层映射到业务层
	modelToimpl.MapFavorite(videos, videos2)

	return videos2, nil
}
