package models

import (
	"time"

	"gorm.io/gorm"
)

type Bid struct {
	ID              uint `gorm:"primaryKey;autoIncrement"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       gorm.DeletedAt `gorm:"index"`
	Name            string         `gorm:"not null"`
	Description     string
	Status          string
	TenderID        uint   `gorm:"references:Tender(id);onDelete:CASCADE"`
	OrganizationID  uint   `gorm:"references:Organization(id);onDelete:CASCADE"`
	Creator         User   `gorm:"foreignKey:CreatorUsername;references:Username"`
	CreatorUsername string `gorm:"references:User(Username);foreignKey:CreatorUsername"`
	Version         int
}

type BidRequest struct {
	Name            string `json:"name" binding:"required"`
	Description     string `json:"description"`
	Status          string `json:"status"`
	TenderID        uint   `json:"tender_id" binding:"required"`
	OrganizationID  uint   `json:"organization_id" binding:"required"`
	CreatorUsername string `json:"creator_username" binding:"required"`
	Version         int    `json:"version"`
}
