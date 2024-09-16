package models

import (
	"time"

	"gorm.io/gorm"
)

type Review struct {
	ID             uint `gorm:"primaryKey;autoIncrement"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      gorm.DeletedAt `gorm:"index"`
	BidID          uint           `gorm:"references:Bid(id);onDelete:CASCADE" json:"bid_id" binding:"required"`
	OrganizationID uint           `gorm:"references:Organization(id);onDelete:CASCADE" json:"organization_id" binding:"required"`
	Review         string         `gorm:"not null" json:"review" binding:"required"`
}

type ReviewRequest struct {
	BidID          uint   `json:"bid_id" binding:"required"`
	OrganizationID uint   `json:"organization_id" binding:"required"`
	Review         string `json:"review" binding:"required"`
}
