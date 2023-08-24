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
