package dao

import (
	"context"
	"github.com/sunflower10086/TikTok/interaction/internal/dao/db"
	"github.com/sunflower10086/TikTok/interaction/internal/dao/models"
)

/**
需要实时查询的有:
1.User:
	关注总数
	粉丝总数
	作品数
	获赞数量
	喜欢数
	是否被关注
2.Video:
	点赞总数
	评论总数
	是否被点赞

对外暴露4个接口
	func RealTimeUser(ctx context.Context, u *models.User) error 获取实时的User信息
	func RealTimeVideo(ctx context.Context, v *models.Video) error 获取实时的Video信息
	func CheckIsFavorite(ctx context.Context, videoID int64, userID int64) (bool, error) 判断关注关系
	func CheckIsFollow(ctx context.Context, userID int64, followerID int64) (bool, error) 判断点赞关系
*/

// 获取实时用户信息
func RealTimeUser(ctx context.Context, u *models.User) error {
	followCount, err := calFollowCount(ctx, u.ID)
	if err != nil {
		return err
	}
	followerCount, err := calFollowerCount(ctx, u.ID)
	if err != nil {
		return err
	}
	workCount, err := calWorkCount(ctx, u.ID)
	if err != nil {
		return err
	}
	totalFavorited, err := calTotalFavorited(ctx, u.ID)
	if err != nil {
		return err
	}
	userFavoriteCount, err := calUserFavoriteCount(ctx, u.ID)
	if err != nil {
		return err
	}

	u.OtherInfo.FollowCount = followCount
	u.OtherInfo.FollowerCount = followerCount
	u.OtherInfo.WorkCount = workCount
	u.OtherInfo.TotalFavorited = totalFavorited
	u.OtherInfo.FavoriteCount = userFavoriteCount

	return nil
}

func RealTimeVideo(ctx context.Context, v *models.Video) error {
	videoFavoriteCount, err := calVideoFavoriteCount(ctx, v.ID)
	if err != nil {
		return err
	}
	commentCount, err := calCommentCount(ctx, v.ID)
	if err != nil {
		return err
	}

	v.FavoriteCount = videoFavoriteCount
	v.CommentCount = commentCount

	return nil
}

// 判断当前视频是否被当前用户点赞
func CheckIsFavorite(ctx context.Context, videoID int64, userID int64) (bool, error) {
	conn := db.GetDB().WithContext(ctx)
	var count int64 = 0

	err := conn.Table("user_favorite").Where("user_id = ? and video_id = ?", userID, videoID).Count(&count).Error
	if err != nil {
		return false, err
	}

	return count == 1, nil
}

// 判断某用户是否被当前用户关注
func CheckIsFollow(ctx context.Context, userID int64, followerID int64) (bool, error) {
	conn := db.GetDB().WithContext(ctx)
	var count int64 = 0

	err := conn.Table("user_follower").Where("user_id = ? and follower_id = ?", userID, followerID).Count(&count).Error
	if err != nil {
		return false, err
	}

	return count == 1, nil
}

// 计算视频的总点赞数
func calVideoFavoriteCount(ctx context.Context, videoID int64) (int64, error) {
	conn := db.GetDB().WithContext(ctx)
	var count int64
	err := conn.Table("user_favorite").Where("video_id = ?", videoID).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

// 计算视频的总评论数
func calCommentCount(ctx context.Context, videoID int64) (int64, error) {
	conn := db.GetDB().WithContext(ctx)
	var count int64
	err := conn.Model(&models.Comment{}).Where("video_id = ?", videoID).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

// 计算用户的关注总数
func calFollowCount(ctx context.Context, followerID int64) (int64, error) {
	conn := db.GetDB().WithContext(ctx)
	var count int64 = 0

	err := conn.Table("user_follower").Where("follower_id = ?", followerID).Count(&count).Error
	if err != nil {
		return 0, err
	}

	return count, nil
}

// 计算用户的粉丝总数
func calFollowerCount(ctx context.Context, userID int64) (int64, error) {
	conn := db.GetDB().WithContext(ctx)
	var count int64 = 0

	err := conn.Table("user_follower").Where("user_id = ?", userID).Count(&count).Error
	if err != nil {
		return 0, err
	}

	return count, nil
}

// 计算用户的作品数
func calWorkCount(ctx context.Context, userID int64) (int64, error) {
	conn := db.GetDB().WithContext(ctx)
	var count int64 = 0

	err := conn.Model(&models.Video{}).Where("author_id = ?", userID).Count(&count).Error
	if err != nil {
		return 0, err
	}

	return count, nil
}

// 计算用户的喜欢数
func calUserFavoriteCount(ctx context.Context, userID int64) (int64, error) {
	conn := db.GetDB().WithContext(ctx)
	var count int64 = 0

	err := conn.Table("user_favorite").Where("user_id = ?", userID).Count(&count).Error
	if err != nil {
		return 0, err
	}

	return count, nil
}

// 计算用户的获赞数量
func calTotalFavorited(ctx context.Context, userID int64) (int64, error) {
	conn := db.GetDB().WithContext(ctx)
	var count int64 = 0
	var videoID []int64

	// 查询用户所有作品的videoID
	err := conn.Find(&models.Video{}).Where("author_id = ?", userID).Pluck("video_id", &videoID).Error
	if err != nil {
		return 0, err
	}

	// 累加每个视频的点赞数
	for _, ID := range videoID {
		c, err := calVideoFavoriteCount(ctx, ID)
		if err != nil {
			return 0, err
		}
		count += c
	}

	return count, err
}
