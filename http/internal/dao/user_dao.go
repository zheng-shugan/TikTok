package dao

import (
	"errors"
	"github.com/sunflower10086/TikTok/http/internal/dao/db"
	"github.com/sunflower10086/TikTok/http/internal/models"
	"gorm.io/gorm"
)

// 基于gorm写增删改查user

// CreateUser 增
func CreateUser(user *models.User) error {
	return db.GetDB().Create(user).Error
}

// DeleteUser 删
func DeleteUser(user *models.User) error {
	return db.GetDB().Delete(user).Error
}

// UpdateUser 改
func UpdateUser(user *models.User) error {
	return db.GetDB().Save(user).Error
}

// GetUserByID 查
func GetUserByID(id uint64) (*models.User, error) {
	var user models.User
	err := db.GetDB().Where("id = ?", id).First(&user).Error
	return &user, err
}

// GetUserByUsername 通过用户名查找用户
func GetUserByUsername(username string) (*models.User, error) {
	var user models.User
	err := db.GetDB().Where("username = ?", username).First(&user).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// 如果找不到用户，返回自定义错误
			return nil, nil
		}
		// 如果发生其他错误，直接返回错误
		return nil, err
	}

	return &user, err
}
