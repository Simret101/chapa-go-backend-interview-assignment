package routes

import (
	"chapa/pkg/config"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(r *gin.RouterGroup, cfg *config.Config, db *gorm.DB) {

	SetupTransactionGroup(r.Group("/transactions"), db)

	SetupTransferGroup(r.Group("/transfers"), db)

}
