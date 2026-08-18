package models

import "time"

type Checklist struct {
	ID        uint      `gorm:"primaryKey"`
	TaskID    uint      `gorm:"not null"`
	Title     string    `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time

	Task  Task             `gorm:"foreignKey:TaskID"`
	Items []ChecklistItem  `gorm:"foreignKey:ChecklistID"`
}