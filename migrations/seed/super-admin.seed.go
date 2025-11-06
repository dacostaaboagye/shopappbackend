package seed

import (
	"errors"
	"log"

	"github.com/Aboagye-Dacosta/shopBackend/internal/database/models"
	"github.com/Aboagye-Dacosta/shopBackend/internal/env"
	"github.com/Aboagye-Dacosta/shopBackend/internal/utils"
	"gorm.io/gorm"
)

func SeedSuperAdmin(db *gorm.DB) error {
	email := env.GetStringEnv("SUPER_ADMIN_EMAIL", "")
	password := env.GetStringEnv("SUPER_ADMIN_PASS", "")

	if email == "" || password == "" {
		return errors.New("SUPER_ADMIN_EMAIL or SUPER_ADMIN_PASS not set in environment")
	}

	var user models.User
	err := db.Where("email = ?", email).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		hashed, err := utils.HashPassword(password)
		if err != nil {
			return err
		}

		user = models.User{
			FirstName: "Super",
			LastName:  "Admin",
			Email:     email,
			Password:  hashed,
		}

		if err := db.Create(&user).Error; err != nil {
			return err
		}

		log.Println("✅ Super Admin created successfully.")
	} else if err != nil {
		return err
	} else {
		log.Println("ℹ️ Super Admin already exists, skipping creation.")
	}

	// Ensure admin role exists
	var role models.Role
	if err := db.Where("name = ?", "admin").First(&role).Error; err != nil {
		return err
	}

	// Attach the admin role if not already assigned
	hasRole := db.Model(&user).Where("role_id = ?", role.ID).Association("Roles").Count() > 0
	if !hasRole {
		if err := db.Model(&user).Association("Roles").Append(&role); err != nil {
			return err
		}
		log.Println("✅ Admin role assigned to Super Admin.")
	}

	return nil
}
