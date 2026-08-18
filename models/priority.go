package models

import "time"

type Priority struct {
	ID uint `gorm:"primaryKey"`
	Name string `gorm:"not null"`
	Level int `gorm:"default:0"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Tasks []Task `gorm:"foreignKey:PriorityID"`
}