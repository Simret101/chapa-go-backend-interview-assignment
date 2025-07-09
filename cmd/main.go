package main

import (
	"log"

	"chapa/api/routes"
	"chapa/migration" // Ensure migration package is imported
	"chapa/pkg/config"
	"chapa/pkg/database"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	
	cfg := config.LoadConfig()

	
	db := database.ConnectDB()

	
	cache := config.NewRedisCache()

	
	err := cache.SetCache("example_key", "example_value")
	if err != nil {
		log.Println("Error setting cache:", err)
	} else {
		log.Println("Cache set successfully")
	}

	val, err := cache.GetCache("example_key")
	if err != nil {
		log.Println("Error getting cache:", err)
	} else {
		log.Println("Cache value:", val)
	}

	err = cache.DeleteCache("example_key")
	if err != nil {
		log.Println("Error deleting cache:", err)
	} else {
		log.Println("Cache deleted successfully")
	}

	
	router := gin.Default()


	router.Use(cors.Default())

	
	if err := migration.MigrateModels(db); err != nil {
		log.Fatalf("Migration failed: %v", err)
	}


	routes.SetupRoutes(router.Group("/api/v0"), &cfg, db)

	
	log.Printf("Server is running on port %s\n", cfg.AppPort)

	if err := router.Run(cfg.AppPort); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
