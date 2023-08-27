package models

import (
	"gorm.io/gorm"
)

type CustomBaseModel struct {
	ID        int64 `gorm:"primaryKey"`
	CreatedAt int64
	UpdatedAt int64
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
