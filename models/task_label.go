package models

import "time"

type TaskLabel struct {
	ID uint `gorm:"primaryKey"`
	TaskID uint `gorm:"not null"`
	LabelID uint `gorm:"not null"`
	CreatedAt time.Time
	Task Task `gorm:"foreignKey:TaskID"`
	Label Label `gorm:"foreignKey:LabelID`
}