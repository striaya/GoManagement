package models

import "time"

type Task struct{
	ID uint `gorm:"primaryKey"`
	ProjectID uint `gorm:"not null"`
	StatusID uint `gorm:"not null"`
	PriorityID uint `gorm:"not null"`
	CreatedByID uint `gorm:"not null"`
	ParentID *uint
	Title string `gorm:"not null"`
	Description string
	TaskNumber int `gorm:"default:0"`
	DueDate *time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
	Project Project `gorm:"foreignKey:ProjectID"`
	Status Status `gorm:"foreignKey:StatusID"`
	Priority Priority `gorm:"foreignKey:PriorityID"`
	CreatedBy User `gorm:"foreignKey:CreatedByID"`
	Parent *Task `gorm:"foreignKey:ParentID"`
	SubTasks []Task `gorm:"foreignKey:ParentID"`
}