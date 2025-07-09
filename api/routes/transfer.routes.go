package routes

import (
	"chapa/api/controllers"
	"chapa/internal/repositories"
	"chapa/internal/usecases"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupTransferGroup(transferRoute *gin.RouterGroup, db *gorm.DB) {
	repo := repositories.NewTransferRepository(db)
	webhookSecret := os.Getenv("CHAPA_WEBHOOK_SECRET")
	uc := usecases.NewTransferUseCase(repo, webhookSecret)
	ctrl := controllers.NewTransferController(uc)

	transferRoute.POST("/initiate", ctrl.Initiate)
	transferRoute.GET("/verify/:ref", ctrl.Verify)
	transferRoute.POST("/webhook", ctrl.Webhook)
}
