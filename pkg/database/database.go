package database

import (
	"chapa/pkg/config"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func ConnectDB() *gorm.DB {
	// Load configuration
	cfg := config.LoadConfig()

	
	dsn := cfg.DBURL

	// Connect to the database
	var err error
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), // Set to Info level for development
	})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Get the underlying SQL DB object
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Failed to get DB object: %v", err)
	}
	// Verify connection
	if err := sqlDB.Ping(); err != nil {
		log.Fatalf("Failed to ping DB: %v", err.Error())
	} else {
		fmt.Println("Connected to ping DB")
	}

	return db
}
