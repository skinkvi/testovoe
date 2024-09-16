package handlers

import (
	"testovoeZadanir/internal/models"
	"testovoeZadanir/internal/util"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func GetReviewsHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		bidId := c.Param("bidId")

		var reviews []models.Review

		db.Where("bid_id = ?", bidId).Find(&reviews)

		util.Logger.Info("Reviews retrieved", zap.String("bid_id", bidId))
		c.JSON(200, reviews)
	}
}
