package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint `gorm:"primaryKey;autoIncrement"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Username  string         `gorm:"unique;not null" json:"username" binding:"required"`
	FirstName string         `json:"first_name"`
	LastName  string         `json:"last_name"`
}
