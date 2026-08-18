package models

import "time"

type Status struct {
	ID        uint      `gorm:"primaryKey"`
	Name      string    `gorm:"not null"`
	Color     string
	Position  int       `gorm:"default:0"`
	CreatedAt time.Time
	UpdatedAt time.Time

	Tasks []Task `gorm:"foreignKey:StatusID"`
}