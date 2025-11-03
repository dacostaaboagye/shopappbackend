package db

import (
	"log"

	"github.com/Aboagye-Dacosta/shopBackend/internal/env"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {
	dsn := env.GetStringEnv("DATABASE_URL", "")

	if dsn == "" {
		log.Fatal("❌ DATABASE_URL not set")
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to Connect to Database", err)
	}

	runMigration(db)

	return db
}
