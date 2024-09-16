package models

import (
	"time"

	"gorm.io/gorm"
)

type Decision struct {
	ID             uint `gorm:"primaryKey;autoIncrement"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      gorm.DeletedAt `gorm:"index"`
	BidID          uint           `gorm:"references:Bid(id);onDelete:CASCADE"`
	OrganizationID uint           `gorm:"references:Organization(id);onDelete:CASCADE"`
	Decision       string
}
