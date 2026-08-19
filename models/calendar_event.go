package models

import "time"

type CalendarEvent struct {
	ID          uint      `gorm:"primaryKey"`
	WorkspaceID uint      `gorm:"not null"`
	CreatedByID uint      `gorm:"not null"`
	Title       string    `gorm:"not null"`
	Description string
	StartAt     time.Time `gorm:"not null"`
	EndAt       time.Time `gorm:"not null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time

	Workspace Workspace `gorm:"foreignKey:WorkspaceID"`
	CreatedBy User      `gorm:"foreignKey:CreatedByID"`
}