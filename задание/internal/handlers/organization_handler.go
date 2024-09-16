package handlers

import (
	"strconv"
	"testovoeZadanir/internal/models"
	"testovoeZadanir/internal/util"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func CreateOrganizationHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var organization models.Organization

		if err := c.ShouldBindJSON(&organization); err != nil {
			util.Logger.Error("Failed to bind JSON", zap.Error(err))
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		util.Logger.Info("Creating organization",
			zap.String("name", organization.Name),
			zap.String("description", organization.Description),
			zap.String("type", organization.Type))

		if err := db.Create(&organization).Error; err != nil {
			util.Logger.Error("Failed to create organization", zap.Error(err))
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		util.Logger.Info("Organization created", zap.String("organization_id", strconv.Itoa(int(organization.ID))))
		c.JSON(200, organization)
	}
}

func CreateOrganizationResponsibleHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var responsible models.OrganizationResponsible

		if err := c.ShouldBindJSON(&responsible); err != nil {
			util.Logger.Error("Failed to bind JSON", zap.Error(err))
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		util.Logger.Info("Creating organization responsible",
			zap.Uint("organization_id", responsible.OrganizationID),
			zap.Uint("user_id", responsible.UserID))

		if err := db.Create(&responsible).Error; err != nil {
			util.Logger.Error("Failed to create organization responsible", zap.Error(err))
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		util.Logger.Info("Organization responsible created", zap.String("responsible_id", strconv.Itoa(int(responsible.ID))))
		c.JSON(200, responsible)
	}
}

func GetOrganizationHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			util.Logger.Error("Invalid organization id", zap.Error(err))
			c.JSON(400, gin.H{"error": "Invalid organization id"})
			return
		}

		var organization models.Organization
		if err := db.First(&organization, id).Error; err != nil {
			util.Logger.Error("Organization not found", zap.Int("organization_id", id), zap.Error(err))
			c.JSON(404, gin.H{"error": "Organization not found"})
			return
		}

		c.JSON(200, organization)
	}
}

func UpdateOrganizationHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			util.Logger.Error("Invalid organization id", zap.Error(err))
			c.JSON(400, gin.H{"error": "Invalid organization id"})
			return
		}

		var organization models.Organization
		if err := db.First(&organization, id).Error; err != nil {
			util.Logger.Error("Organization not found", zap.Int("organization_id", id), zap.Error(err))
			c.JSON(404, gin.H{"error": "Organization not found"})
			return
		}

		var updateOrganization models.Organization
		if err := c.ShouldBindJSON(&updateOrganization); err != nil {
			util.Logger.Error("Failed to bind JSON", zap.Error(err))
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		if updateOrganization.Name != "" {
			organization.Name = updateOrganization.Name
		}
		if updateOrganization.Description != "" {
			organization.Description = updateOrganization.Description
		}
		if updateOrganization.Type != "" {
			organization.Type = updateOrganization.Type
		}

		db.Save(&organization)

		util.Logger.Info("Organization updated", zap.Int("organization_id", id))
		c.JSON(200, organization)
	}
}

func DeleteOrganizationHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			util.Logger.Error("Invalid organization id", zap.Error(err))
			c.JSON(400, gin.H{"error": "Invalid organization id"})
			return
		}

		if err := db.Delete(&models.Organization{}, id).Error; err != nil {
			util.Logger.Error("Failed to delete organization", zap.Int("organization_id", id), zap.Error(err))
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		util.Logger.Info("Organization deleted", zap.Int("organization_id", id))
		c.Status(204)
	}
}
