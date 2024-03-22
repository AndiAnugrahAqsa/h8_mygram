package models

import "time"

type SocialMedia struct {
	ID             int
	Name           string `gorm:"not null;type:varchar(150)"`
	SocialMediaURL string `gorm:"not null"`
	UserID         int    `gorm:"not null"`
	User           User
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
