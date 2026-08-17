package models 

import (
	"time"
	"gorm.io/gorm"
)

type Workspace struct {
	ID uint `gorm:"primaryKey"`
	Name string `gorm:"not null"`
	Description string
	OwnerID uint `gorm:"not null"`
	Owner User `gorm:"foreignKey:OwnerID"`
	CreatedAt time.time
	UPdatedAt time.time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}