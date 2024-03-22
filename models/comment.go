package models

import "time"

type Comment struct {
	ID        int
	UserID    int `gorm:"not null"`
	User      User
	PhotoID   int `gorm:"not null"`
	Photo     Photo
	Message   string `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
