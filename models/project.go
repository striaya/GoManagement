package models

import "time"

type Project struct {
	ID uint `gorm:"primaryKey"`
	WorkspaceID uint `gorm:"not null"`
	OwnerID uint `gorm:"not null"`
	Name string `gorm:"not null"`
	Key string `gorm:"not null"`
	Description string
	Status string `gorm:"not null;default:'active'`
	CreatedAt time.Time
	UpdatedAt time.Time
	Workspace Workspace `gorm:"foreignKey:WorkspaceID"`
	Owner User `gorm:"foreignKey:OwnerID"`
}