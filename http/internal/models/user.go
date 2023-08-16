package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username        string `gorm:"type:varchar(20);not null;unique" json:"username" binding:"required"`
	Password        string `gorm:"type:varchar(40);not null" json:"password" binding:"required"`
	Avatar          string `gorm:"type:varchar(255)" json:"avatar"`
	BackgroundImage string `gorm:"type:varchar(255)" json:"background_image"`
	Signature       string `gorm:"type:varchar(255)" json:"signature"`
}
