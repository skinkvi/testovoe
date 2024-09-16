package handlers

import (
	"strconv"
	"testovoeZadanir/internal/models"
	"testovoeZadanir/internal/util"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func CreateReviewHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var reviewRequest models.ReviewRequest

		if err := c.ShouldBindJSON(&reviewRequest); err != nil {
			util.Logger.Error("Failed to bind JSON", zap.Error(err))
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		var review models.Review
		review.BidID = reviewRequest.BidID
		review.OrganizationID = reviewRequest.OrganizationID
		review.Review = reviewRequest.Review

		if err := db.Create(&review).Error; err != nil {
			util.Logger.Error("Failed to create review", zap.String("review_id", strconv.FormatUint(uint64(review.ID), 10)), zap.Error(err))
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		util.Logger.Info("Review created", zap.String("review_id", strconv.FormatUint(uint64(review.ID), 10)))
		c.JSON(200, review)
	}
}
