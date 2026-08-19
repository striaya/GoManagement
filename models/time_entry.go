package models

import "time"

type TimeEntry struct {
	ID        uint      `gorm:"primaryKey"`
	TaskID    uint      `gorm:"not null"`
	UserID    uint      `gorm:"not null"`
	StartedAt time.Time `gorm:"not null"`
	EndedAt   *time.Time
	Duration  int       `gorm:"default:0"`
	Note      string
	CreatedAt time.Time
	UpdatedAt time.Time

	Task Task `gorm:"foreignKey:TaskID"`
	User User `gorm:"foreignKey:UserID"`
}