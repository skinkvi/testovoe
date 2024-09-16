package main

import (
	"os"
	"testovoeZadanir/internal/handlers"
	"testovoeZadanir/internal/models"
	"testovoeZadanir/internal/util"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func main() {
	err := godotenv.Load()
	if err != nil {
		util.Logger.Fatal("Error loading .env file", zap.Error(err))
	}

	dsn := os.Getenv("POSTGRES_CONN")
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		util.Logger.Fatal("failed to connect database", zap.Error(err))
	}

	db.AutoMigrate(&models.User{}, &models.Organization{}, &models.OrganizationResponsible{}, &models.Tender{}, &models.Bid{}, &models.Decision{}, &models.Review{})

	r := gin.Default()

	r.GET("/ping", handlers.PingHandler)

	r.POST("/api/tenders/new", handlers.CreateTenderHandler(db))
	r.GET("/api/tenders", handlers.GetTendersHandler(db))
	r.GET("/api/tenders/my", handlers.GetMyTendersHandler(db))
	r.PATCH("/api/tenders/:tenderId/edit", handlers.EditTenderHandler(db))
	r.PUT("/api/tenders/:tenderId/rollback/:version", handlers.RollbackTenderHandler(db))

	r.POST("/api/bids/new", handlers.CreateBidHandler(db))
	r.GET("/api/bids/my", handlers.GetMyBidsHandler(db))
	r.GET("/api/bids/tender/:tenderId/list", handlers.GetBidsForTenderHandler(db))
	r.PATCH("/api/bids/:bidId/edit", handlers.EditBidHandler(db))
	r.PUT("/api/bids/:bidId/rollback/:version", handlers.RollbackBidHandler(db))

	r.POST("/api/decisions/new", handlers.CreateDecisionHandler(db))
	r.GET("/api/bids/approve/:bidId", handlers.ApproveBidHandler(db))

	r.POST("/api/reviews/new", handlers.CreateReviewHandler(db))
	r.GET("/api/bids/:bidId/reviews", handlers.GetReviewsHandler(db))

	r.POST("/api/users/new", handlers.CreateUserHandler(db))
	r.GET("/api/users", handlers.GetUserByIdHandler(db))
	r.POST("/api/users/:id/edit", handlers.UpdateUserHandler(db))
	r.DELETE("/api/users/:id/delete", handlers.DeleteUserHandler(db))

	r.POST("/api/organizations/new", handlers.CreateOrganizationHandler(db))
	r.GET("/api/organizations", handlers.GetOrganizationHandler(db))
	r.POST("/api/organizations/:id/responsible/edit", handlers.UpdateOrganizationHandler(db))
	r.GET("/api/organizations/:id/bids", handlers.DeleteOrganizationHandler(db))
	r.POST("/api/organizations/responsible", handlers.CreateOrganizationResponsibleHandler(db))

	serverAddress := os.Getenv("SERVER_ADDRESS")
	if serverAddress == "" {
		serverAddress = "0.0.0.0:8080"
	}

	r.Run(serverAddress)
}
