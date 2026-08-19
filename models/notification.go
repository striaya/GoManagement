package models

import "time"

type Notification struct {
	ID        uint      `gorm:"primaryKey"`
	UserID    uint      `gorm:"not null"`
	Title     string    `gorm:"not null"`
	Message   string    `gorm:"not null"`
	Type      string
	IsRead    bool      `gorm:"default:false"`
	CreatedAt time.Time

	User User `gorm:"foreignKey:UserID"`
}