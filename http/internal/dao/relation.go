package dao

import (
	"github.com/sunflower10086/TikTok/http/internal/dao/db"
	"github.com/sunflower10086/TikTok/http/internal/models"
	"log"
)

// FindRelation 给定当前用户和目标用户id，查询 follow 表中相应的记录
func FindRelation(userId int64, followerId int64) (*models.Follow, error) {
	follow := models.Follow{}

	err := db.GetDB().
		Where("user_id = ?", userId).
		Where("follower_id = ?", followerId).
		Take(&follow).Error

	if err != nil {
		// 没有记录不视为错误
		if "record not found" == err.Error() {
			return nil, nil
		}
		log.Panicln(err.Error())
	}
	return &follow, nil
}

// GetFollowerIds 根据用户 id 获取粉丝的 id 列表
func GetFollowerIds(userId int64) ([]int64, error) {
	idList := make([]int64, 0)

	err := db.GetDB().
		Where("user_id = ?", userId).
		Take(&idList).Error

	if err != nil {
		if "record not found" == err.Error() {
			return nil, nil
		}
		log.Panicln(err.Error())
	}

	return idList, nil
}

// GetFollowerList 根据用户 id 获取粉丝列表
func GetFollowerList(userId int64) ([]*models.User, error) {
	followerList := make([]*models.User, 0)
	followerIds, _ := GetFollowerIds(userId)

	err := db.GetDB().
		Where("user_id IN ?", followerIds).
		Take(&followerList).Error

	if err != nil {
		log.Panicln(err.Error())
	}

	return followerList, nil
}
