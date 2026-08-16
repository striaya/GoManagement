package models 

import (
	"time"
	"gorm.io/gorm"
)

type User struct {
	ID uint `gorm:"primaryKey"`
	Name string `gorm:"not null"`
	Email string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
	Avatar string
	RoleID uint `gorm:"not null"`
	Role Role `gorm:"foreignKey:RoleID"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

}