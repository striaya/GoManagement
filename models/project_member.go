package models

import "time"

type ProjectMember struct {
	ID        uint      `gorm:"primaryKey"`
	ProjectID uint      `gorm:"not null"`
	UserID    uint      `gorm:"not null"`
	Role      string    `gorm:"not null;default:'member'"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Project Project `gorm:"foreignKey:ProjectID"`
	User    User    `gorm:"foreignKey:UserID"`
}