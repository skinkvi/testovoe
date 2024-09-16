package handlers

import (
	"errors"
	"strconv"
	"testovoeZadanir/internal/models"
	"testovoeZadanir/internal/util"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func CreateBidHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var bidRequest models.BidRequest

		if err := c.ShouldBindJSON(&bidRequest); err != nil {
			util.Logger.Error("Failed to bind JSON", zap.Error(err))
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		var bid models.Bid
		bid.Name = bidRequest.Name
		bid.Description = bidRequest.Description
		bid.Status = bidRequest.Status
		bid.TenderID = bidRequest.TenderID
		bid.OrganizationID = bidRequest.OrganizationID
		bid.CreatorUsername = bidRequest.CreatorUsername
		bid.Version = bidRequest.Version

		if err := db.Create(&bid).Error; err != nil {
			util.Logger.Error("Failed to create bid", zap.Error(err))
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		util.Logger.Info("Bid created", zap.String("bid_id", strconv.FormatUint(uint64(bid.ID), 10)))
		c.JSON(200, bid)
	}
}

func GetMyBidsHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		username := c.Query("username")

		if username == "" {
			util.Logger.Error("Username is empty", zap.String("username", username))
			c.JSON(400, gin.H{"error": "Username is empty"})
			return
		}

		var bids []models.Bid

		if err := db.Where("creator_username = ?", username).Find(&bids).Error; err != nil {
			util.Logger.Error("Failed to retrieve bids for user", zap.String("username", username), zap.Error(err))
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		if bids == nil {
			util.Logger.Error("Bids is nil", zap.String("username", username))
			c.JSON(500, gin.H{"error": "Bids is nil"})
			return
		}

		util.Logger.Info("My bids retrieved", zap.String("username", username), zap.Int("count", len(bids)))
		c.JSON(200, bids)
	}
}

func GetBidsForTenderHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		tenderId := c.Param("tenderId")

		if db == nil {
			util.Logger.Error("Database connection is nil")
			c.JSON(500, gin.H{"error": "Database connection is nil"})
			return
		}

		var bids []models.Bid

		if err := db.Where("tender_id = ?", tenderId).Find(&bids).Error; err != nil {
			util.Logger.Error("Failed to retrieve bids for tender", zap.String("tender_id", tenderId), zap.Error(err))
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		if bids == nil {
			util.Logger.Error("Bids is nil", zap.String("tender_id", tenderId))
			c.JSON(500, gin.H{"error": "Bids is nil"})
			return
		}

		util.Logger.Info("Bids for tender retrieved", zap.String("tender_id", tenderId), zap.Int("count", len(bids)))
		c.JSON(200, bids)
	}
}

func EditBidHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		bidId := c.Param("bidId")

		var bid models.Bid
		if err := db.First(&bid, bidId).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				util.Logger.Error("Bid not found", zap.String("bid_id", bidId), zap.Error(err))
				c.JSON(404, gin.H{"error": "Bid not found"})
				return
			}

			util.Logger.Error("Failed to retrieve bid", zap.String("bid_id", bidId), zap.Error(err))
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		var bidRequest models.BidRequest
		if err := c.ShouldBindJSON(&bidRequest); err != nil {
			util.Logger.Error("Failed to bind JSON", zap.Error(err))
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		if bidRequest.Name != "" {
			bid.Name = bidRequest.Name
		}
		if bidRequest.Description != "" {
			bid.Description = bidRequest.Description
		}
		if bidRequest.Status != "" {
			bid.Status = bidRequest.Status
		}
		if bidRequest.TenderID != 0 {
			bid.TenderID = bidRequest.TenderID
		}
		if bidRequest.OrganizationID != 0 {
			bid.OrganizationID = bidRequest.OrganizationID
		}
		if bidRequest.CreatorUsername != "" {
			bid.CreatorUsername = bidRequest.CreatorUsername
		}
		bid.Version++

		if err := db.Save(&bid).Error; err != nil {
			util.Logger.Error("Failed to save bid", zap.String("bid_id", bidId), zap.Error(err))
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		util.Logger.Info("Bid edited", zap.String("bid_id", bidId))
		c.JSON(200, bid)
	}
}

func RollbackBidHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		bidId := c.Param("bidId")
		version, err := strconv.Atoi(c.Param("version"))
		if err != nil {
			util.Logger.Error("Invalid version", zap.String("bid_id", bidId), zap.Error(err))
			c.JSON(400, gin.H{"error": "Invalid version"})
			return
		}

		var bid models.Bid
		if err := db.First(&bid, bidId).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				util.Logger.Error("Bid not found", zap.String("bid_id", bidId), zap.Error(err))
				c.JSON(404, gin.H{"error": "Bid not found"})
				return
			}

			util.Logger.Error("Failed to retrieve bid", zap.String("bid_id", bidId), zap.Error(err))
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		if bid.Version < version {
			util.Logger.Error("Version is newer than current", zap.String("bid_id", bidId), zap.Int("current_version", bid.Version), zap.Int("requested_version", version))
			c.JSON(409, gin.H{"error": "Version is newer than current"})
			return
		}

		if bid.Version > version {
			bid.Version--
			if err := db.Save(&bid).Error; err != nil {
				util.Logger.Error("Failed to save bid", zap.String("bid_id", bidId), zap.Error(err))
				c.JSON(500, gin.H{"error": err.Error()})
				return
			}
		}

		util.Logger.Info("Bid rolled back", zap.String("bid_id", bidId), zap.Int("version", version))
		c.JSON(200, bid)
	}
}
