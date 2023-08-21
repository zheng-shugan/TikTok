package models

import (
	"gorm.io/gorm"
	"time"
)

type CustomBaseModel struct {
	ID        int64 `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
