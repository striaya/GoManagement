package models

import "time"

type Attachment struct {
	ID        uint      `gorm:"primaryKey"`
	TaskID    uint      `gorm:"not null"`
	UserID    uint      `gorm:"not null"`
	FileName  string    `gorm:"not null"`
	FilePath  string    `gorm:"not null"`
	FileType  string
	FileSize  int64
	CreatedAt time.Time

	Task Task `gorm:"foreignKey:TaskID"`
	User User `gorm:"foreignKey:UserID"`
}