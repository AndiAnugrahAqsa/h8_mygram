package models

import "time"

type Photo struct {
	ID        int
	Title     string `gorm:"not null;type:varchar(255)"`
	Caption   string
	PhotoURL  string `gorm:"not null"`
	UserID    int    `gorm:"not null"`
	User      User
	Comments  []Comment
	CreatedAt time.Time
	UpdatedAt time.Time
}
