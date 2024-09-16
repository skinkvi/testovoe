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

func CreateUserHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user models.User

		if err := c.ShouldBindJSON(&user); err != nil {
			util.Logger.Error("Failed to bind JSON", zap.Error(err))
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		if user == (models.User{}) {
			util.Logger.Error("User is empty")
			c.JSON(400, gin.H{"error": "User is empty"})
			return
		}

		if err := db.Create(&user).Error; err != nil {
			util.Logger.Error("Failed to create user", zap.Error(err))
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		util.Logger.Info("User created", zap.String("user_id", strconv.Itoa(int(user.ID))))
		c.JSON(200, user)
	}
}

func GetUserByIdHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			util.Logger.Error("Invalid user id", zap.Error(err))
			c.JSON(400, gin.H{"error": "Invalid user id"})
			return
		}

		var user models.User
		if err := db.First(&user, id).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				util.Logger.Error("User not found", zap.Int("user_id", id), zap.Error(err))
				c.JSON(404, gin.H{"error": "User not found"})
				return
			}

			util.Logger.Error("Failed to retrieve user", zap.Int("user_id", id), zap.Error(err))
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		if user.ID == 0 {
			util.Logger.Error("User ID is zero", zap.Int("user_id", id))
			c.JSON(400, gin.H{"error": "User ID is zero"})
			return
		}

		c.JSON(200, user)
	}
}

func UpdateUserHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			util.Logger.Error("Invalid user id", zap.Error(err))
			c.JSON(400, gin.H{"error": "Invalid user id"})
			return
		}

		var user models.User
		if err := db.First(&user, id).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				util.Logger.Error("User not found", zap.Int("user_id", id), zap.Error(err))
				c.JSON(404, gin.H{"error": "User not found"})
				return
			}

			util.Logger.Error("Failed to retrieve user", zap.Int("user_id", id), zap.Error(err))
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		if user.ID == 0 {
			util.Logger.Error("User ID is zero", zap.Int("user_id", id))
			c.JSON(400, gin.H{"error": "User ID is zero"})
			return
		}

		var updateUser models.User
		if err := c.ShouldBindJSON(&updateUser); err != nil {
			util.Logger.Error("Failed to bind JSON", zap.Error(err))
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		if updateUser.Username != "" {
			user.Username = updateUser.Username
		}
		if updateUser.FirstName != "" {
			user.FirstName = updateUser.FirstName
		}
		if updateUser.LastName != "" {
			user.LastName = updateUser.LastName
		}

		if err := db.Save(&user).Error; err != nil {
			util.Logger.Error("Failed to save user", zap.Int("user_id", id), zap.Error(err))
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		util.Logger.Info("User updated", zap.Int("user_id", id))
		c.JSON(200, user)
	}
}

func DeleteUserHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			util.Logger.Error("Invalid user id", zap.Error(err))
			c.JSON(400, gin.H{"error": "Invalid user id"})
			return
		}

		var user models.User
		if err := db.First(&user, id).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				util.Logger.Error("User not found", zap.Int("user_id", id), zap.Error(err))
				c.JSON(404, gin.H{"error": "User not found"})
				return
			}

			util.Logger.Error("Failed to retrieve user", zap.Int("user_id", id), zap.Error(err))
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		if err := db.Delete(&user).Error; err != nil {
			util.Logger.Error("Failed to delete user", zap.Int("user_id", id), zap.Error(err))
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		util.Logger.Info("User deleted", zap.Int("user_id", id))
		c.Status(204)
	}
}
