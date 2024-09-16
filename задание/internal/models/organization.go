package models

import (
	"time"

	"gorm.io/gorm"
)

type Organization struct {
	ID          uint `gorm:"primaryKey;autoIncrement"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
	Name        string         `gorm:"not null" json:"name" binding:"required"`
	Description string         `json:"description"`
	Type        string         `json:"type"`
}

type OrganizationResponsible struct {
	ID             uint `gorm:"primaryKey;autoIncrement"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      gorm.DeletedAt `gorm:"index"`
	OrganizationID uint           `gorm:"references:Organization(id);onDelete:CASCADE" json:"organization_id"`
	UserID         uint           `gorm:"references:User(id);onDelete:CASCADE" json:"user_id"`
}
