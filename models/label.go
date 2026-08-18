package models

import "time"

type Label struct {
	ID uint `gorm:"primaryKey"`
	ProjectID uint `gorm:"not null"`
	Name string `gorm:"not null"`
	Color string
	CreatedAt time.Time
	UpdatedAt time.Time
	Project Project `gorm:"foreignKey:ProjectID"`
}