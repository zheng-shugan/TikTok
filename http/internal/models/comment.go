package models

type Comment struct {
	CustomBaseModel
	VideoID int64  `gorm:"not null"`
	Video   Video  `gorm:"foreignKey:VideoID"`
	UserID  int64  `gorm:"not null"`
	User    User   `gorm:"foreignKey:UserID"`
	Content string `gorm:"type:varchar(255);not null" json:"content"`
}
