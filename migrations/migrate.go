package main

import (
	"log"

	database "github.com/Aboagye-Dacosta/shopBackend/internal/database/db"
	"github.com/Aboagye-Dacosta/shopBackend/internal/database/models"
	"github.com/Aboagye-Dacosta/shopBackend/internal/env"
	"github.com/Aboagye-Dacosta/shopBackend/migrations/seed"
	"gorm.io/gorm"
)

func main() {
	env.LoadEnv()
	db := database.ConnectDB()

	appModels := []interface{}{
		&models.Permission{},
		&models.Role{},
		&models.User{},
		&models.Customer{},
		&models.Product{},
		&models.Order{},
		&models.Payment{},
	}

	if err := migrateAndSeed(db, appModels...); err != nil {
		log.Fatalf("❌ Migration failed: %v", err)
	}
	log.Println("✅ Migration and seeding completed successfully!")
}

func migrateAndSeed(db *gorm.DB, models ...interface{}) error {
	log.Println("🚀 Running database migrations...")
	if err := db.AutoMigrate(models...); err != nil {
		return err
	}

	log.Println("🌱 Seeding initial data...")
	if err := seed.SeedPermissions(db); err != nil {
		return err
	}
	if err := seed.SeedAdminRole(db); err != nil {
		return err
	}
	if err := seed.SeedUserRole(db); err != nil {
		return err
	}
	if err := seed.SeedSuperAdmin(db); err != nil {
		return err
	}

	return nil
}
