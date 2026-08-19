package models

import "time"

type Activity struct {
	ID         uint      `gorm:"primaryKey"`
	UserID     uint      `gorm:"not null"`
	ProjectID  *uint
	TaskID     *uint
	Action     string    `gorm:"not null"`
	Description string
	CreatedAt  time.Time

	User    User     `gorm:"foreignKey:UserID"`
	Project *Project `gorm:"foreignKey:ProjectID"`
	Task    *Task    `gorm:"foreignKey:TaskID"`
}