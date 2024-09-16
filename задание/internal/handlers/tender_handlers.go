package handlers

import (
	"strconv"
	"testovoeZadanir/internal/models"
	"testovoeZadanir/internal/util"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func CreateTenderHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		util.Logger.Info("Creating tender")

		var tenderRequest models.TenderRequest

		if err := c.ShouldBindJSON(&tenderRequest); err != nil {
			util.Logger.Error("Failed to bind JSON", zap.Error(err))
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		var tender models.Tender
		tender.Name = tenderRequest.Name
		tender.Description = tenderRequest.Description
		tender.ServiceType = tenderRequest.ServiceType
		tender.OrganizationID = tenderRequest.OrganizationID
		tender.CreatorID = tenderRequest.CreatorID
		tender.Version = tenderRequest.Version
		tender.Status = "active"

		if err := db.Create(&tender).Error; err != nil {
			util.Logger.Error("Failed to create tender", zap.Error(err))
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		util.Logger.Info("Tender created", zap.String("tender_id", strconv.FormatUint(uint64(tender.ID), 10)))
		c.JSON(200, tender)
	}
}
func GetTendersHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		util.Logger.Info("Retrieving tenders")

		var tenders []models.Tender

		if db == nil {
			util.Logger.Error("Database connection is nil")
			c.JSON(500, gin.H{"error": "Database connection is nil"})
			return
		}

		if err := db.Find(&tenders).Error; err != nil {
			util.Logger.Error("Failed to retrieve tenders", zap.Error(err))
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		if tenders == nil {
			util.Logger.Error("Tenders is nil")
			c.JSON(500, gin.H{"error": "Tenders is nil"})
			return
		}

		util.Logger.Info("Tenders retrieved", zap.Int("count", len(tenders)))
		c.JSON(200, tenders)
	}
}

func GetMyTendersHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		username := c.Query("username")

		if db == nil {
			util.Logger.Error("Database connection is nil")
			c.JSON(500, gin.H{"error": "Database connection is nil"})
			return
		}

		if username == "" {
			util.Logger.Error("Username is empty", zap.String("username", username))
			c.JSON(400, gin.H{"error": "Username is empty"})
			return
		}

		util.Logger.Info("Retrieving tenders for user", zap.String("username", username))

		var tenders []models.Tender

		if err := db.Joins("JOIN users ON users.id = tenders.creator_id").Where("users.username = ?", username).Find(&tenders).Error; err != nil {
			util.Logger.Error("Failed to retrieve tenders for user", zap.String("username", username), zap.Error(err))
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		if tenders == nil {
			util.Logger.Error("Tenders is nil", zap.String("username", username))
			c.JSON(500, gin.H{"error": "Tenders is nil"})
			return
		}

		util.Logger.Info("My tenders retrieved", zap.String("username", username), zap.Int("count", len(tenders)))
		c.JSON(200, tenders)
	}
}

func EditTenderHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		tenderId := c.Param("tenderId")

		var tender models.Tender
		if err := db.First(&tender, tenderId).Error; err != nil {
			util.Logger.Error("Tender not found", zap.String("tender_id", tenderId), zap.Error(err))
			c.JSON(404, gin.H{"error": "Tender not found"})
			return
		}

		util.Logger.Info("Tender found", zap.String("tender_id", tenderId), zap.String("tender_name", tender.Name))

		var tenderUpdateRequest models.TenderRequest
		if err := c.ShouldBindJSON(&tenderUpdateRequest); err != nil {
			util.Logger.Error("Failed to bind JSON", zap.Error(err))
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		util.Logger.Info("JSON bound", zap.String("tender_id", tenderId), zap.String("tender_name", tenderUpdateRequest.Name))

		if tenderUpdateRequest.Name != "" {
			tender.Name = tenderUpdateRequest.Name
		}
		if tenderUpdateRequest.Description != "" {
			tender.Description = tenderUpdateRequest.Description
		}
		if tenderUpdateRequest.ServiceType != "" {
			tender.ServiceType = tenderUpdateRequest.ServiceType
		}
		if tenderUpdateRequest.OrganizationID != 0 {
			tender.OrganizationID = tenderUpdateRequest.OrganizationID
		}
		if tenderUpdateRequest.CreatorID != 0 {
			tender.CreatorID = tenderUpdateRequest.CreatorID
		}
		tender.Version++

		util.Logger.Info("Tender version incremented", zap.String("tender_id", tenderId), zap.Int("version", tender.Version))

		if err := db.Save(&tender).Error; err != nil {
			util.Logger.Error("Failed to save tender", zap.String("tender_id", tenderId), zap.Error(err))
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		util.Logger.Info("Tender edited", zap.String("tender_id", tenderId))
		c.JSON(200, tender)
	}
}

func RollbackTenderHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		tenderId := c.Param("tenderId")
		version, err := strconv.Atoi(c.Param("version"))
		if err != nil {
			util.Logger.Error("Invalid version", zap.String("tender_id", tenderId), zap.Error(err))
			c.JSON(400, gin.H{"error": "Invalid version"})
			return
		}

		util.Logger.Info("Getting tender", zap.String("tender_id", tenderId))
		var tender models.Tender
		if err := db.First(&tender, tenderId).Error; err != nil {
			util.Logger.Error("Tender not found", zap.String("tender_id", tenderId), zap.Error(err))
			c.JSON(404, gin.H{"error": "Tender not found"})
			return
		}

		util.Logger.Info("Tender found", zap.String("tender_id", tenderId), zap.Int("current_version", tender.Version), zap.Int("requested_version", version))
		if tender.Version < version {
			util.Logger.Error("Version is newer than current", zap.String("tender_id", tenderId), zap.Int("current_version", tender.Version), zap.Int("requested_version", version))
			c.JSON(409, gin.H{"error": "Version is newer than current"})
			return
		}

		if tender.Version > version {
			util.Logger.Info("Rolling back tender version", zap.String("tender_id", tenderId), zap.Int("current_version", tender.Version), zap.Int("requested_version", version))
			tender.Version--
			if err := db.Save(&tender).Error; err != nil {
				util.Logger.Error("Failed to save tender", zap.String("tender_id", tenderId), zap.Error(err))
				c.JSON(500, gin.H{"error": err.Error()})
				return
			}
		}

		util.Logger.Info("Tender rolled back", zap.String("tender_id", tenderId), zap.Int("version", version))
		c.JSON(200, tender)
	}
}
