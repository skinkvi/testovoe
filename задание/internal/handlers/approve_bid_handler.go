package handlers

import (
	"testovoeZadanir/internal/models"
	"testovoeZadanir/internal/util"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func ApproveBidHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		bidId := c.Param("bidId")

		var bid models.Bid
		if err := db.First(&bid, bidId).Error; err != nil {
			util.Logger.Error("Bid not found", zap.String("bid_id", bidId), zap.Error(err))
			c.JSON(404, gin.H{"error": "Bid not found"})
			return
		}

		var decisions []models.Decision
		db.Where("bid_id = ?", bidId).Find(&decisions)

		quorum := min(3, len(decisions))
		approveCount := 0
		rejectCount := 0

		for _, decision := range decisions {
			if decision.Decision == "approve" {
				approveCount++
			} else if decision.Decision == "reject" {
				rejectCount++
			}
		}

		if rejectCount > 0 {
			util.Logger.Info("Bid rejected", zap.String("bid_id", bidId))
			c.JSON(200, gin.H{"status": "rejected"})
			return
		}

		if approveCount >= quorum {
			util.Logger.Info("Bid approved", zap.String("bid_id", bidId))
			c.JSON(200, gin.H{"status": "approved"})
			return
		}

		util.Logger.Info("Bid pending", zap.String("bid_id", bidId))
		c.JSON(200, gin.H{"status": "pending"})
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
