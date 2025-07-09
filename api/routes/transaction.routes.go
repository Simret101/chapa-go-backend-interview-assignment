package routes

import (
    "chapa/api/controllers"
    "chapa/internal/repositories"
    "chapa/internal/usecases"
    "os"

    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
)

func SetupTransactionGroup(txRoute *gin.RouterGroup, db *gorm.DB) {
	repo := repositories.NewTransactionRepository(db)
	webhookSecret := os.Getenv("CHAPA_WEBHOOK_SECRET") 
	uc := usecases.NewTransactionUseCase(repo, webhookSecret)
	ctrl := controllers.NewTransactionController(uc)

	txRoute.POST("/initiate", ctrl.Initiate)
	txRoute.GET("/verify/:txnRef", ctrl.Verify)
	txRoute.POST("/webhook", ctrl.Webhook)
}