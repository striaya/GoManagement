package models

import "time"

type WorkspaceMember struct {
	ID uint `gorm:"primaryKey"`
	WorkspaceID uint `gorm:"not null"`
	UserID uint `gorm:"not null"`
	Role string `gorm:"not null;default:'member'"`
	CreatedAt time.Time
	Workspace Workspace `gorm:"foreignKey:WorkspaceID"`
	User User `gorm:"foreignKey:UserID"`
}
