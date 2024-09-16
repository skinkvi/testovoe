package handlers

import (
	"strconv"
	"testovoeZadanir/internal/models"
	"testovoeZadanir/internal/util"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func CreateDecisionHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var decision models.Decision

		if err := c.ShouldBindJSON(&decision); err != nil {
			util.Logger.Error("Failed to bind JSON", zap.Error(err))
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		if decision == (models.Decision{}) {
			util.Logger.Error("Decision is empty")
			c.JSON(400, gin.H{"error": "Decision is empty"})
			return
		}

		if err := db.Create(&decision).Error; err != nil {
			util.Logger.Error("Failed to create decision", zap.Error(err))
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		util.Logger.Info("Decision created", zap.String("decision_id", strconv.FormatUint(uint64(decision.ID), 10)))
		c.JSON(200, decision)
	}
}
