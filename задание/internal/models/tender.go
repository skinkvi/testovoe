package models

import (
	"time"

	"gorm.io/gorm"
)

type Tender struct {
	ID             uint `gorm:"primaryKey;autoIncrement"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      gorm.DeletedAt `gorm:"index"`
	Name           string         `gorm:"not null" json:"name" binding:"required"`
	Description    string         `json:"description"`
	ServiceType    string         `json:"service_type"`
	Status         string         `json:"status"`
	OrganizationID uint           `gorm:"references:Organization(id);onDelete:CASCADE" json:"organization_id" binding:"required"`
	CreatorID      uint           `gorm:"references:User(id);onDelete:CASCADE" json:"creator_id" binding:"required"`
	Creator        User           `gorm:"foreignKey:CreatorID" json:"creator"`
	Version        int            `json:"version"`
}

type TenderRequest struct {
	Name           string `json:"name" binding:"required"`
	Description    string `json:"description"`
	ServiceType    string `json:"service_type"`
	OrganizationID uint   `json:"organization_id"`
	CreatorID      uint   `json:"creator_id"`
	Version        int    `json:"version"`
}
