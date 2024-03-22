package models

import (
	"time"
)

type User struct {
	ID           int
	Username     string `gorm:"uniqueIndex:idx_username;not null;type:varchar(50)"`
	Email        string `gorm:"uniqueIndex:idx_email;not null;type:varchar(100)"`
	Password     string `gorm:"not null;type:varchar(100)"`
	Age          int    `gorm:"not null"`
	Photos       []Photo
	Comments     []Comment
	SocialMedias []SocialMedia
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
