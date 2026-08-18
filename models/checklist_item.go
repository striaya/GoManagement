package models

import "time"

type ChecklistItem struct {
	ID          uint      `gorm:"primaryKey"`
	ChecklistID uint      `gorm:"not null"`
	Content     string    `gorm:"not null"`
	IsCompleted bool      `gorm:"default:false"`
	Position    int       `gorm:"default:0"`
	CreatedAt   time.Time
	UpdatedAt   time.Time

	Checklist Checklist `gorm:"foreignKey:ChecklistID"`
}