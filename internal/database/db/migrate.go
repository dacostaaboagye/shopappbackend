package db

import (
	"github.com/Aboagye-Dacosta/shopBackend/internal/database/models"
	"gorm.io/gorm"
)

func runMigration(db *gorm.DB) {

	appModels := []interface{}{
		&models.User{},
		&models.Order{},
		&models.Payment{},
		&models.Product{},
	}

	for _, model := range appModels {
		if !db.Migrator().HasTable(model) {
			db.Migrator().CreateTable(model)
		}
	}

}
