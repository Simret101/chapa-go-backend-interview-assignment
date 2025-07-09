package migration

import (
	"chapa/internal/domain"

	"gorm.io/gorm"
)

func MigrateModels(db *gorm.DB) error {
	return db.AutoMigrate(
		&domain.Transaction{},
		&domain.Transfer{},
	)
}
