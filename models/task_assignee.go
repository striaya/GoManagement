package models

import "time"

type TaskAssignee struct {
	ID uint `gorm:"primaryKey"`
	TaskID uint `gorm:"not null"`
	UserID uint `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Task Task `gorm:"foreignKey:TaskID"`
	User User `gorm:"foreignKey:UserID"`
}