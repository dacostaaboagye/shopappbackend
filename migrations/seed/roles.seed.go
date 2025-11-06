package seed

import (
	"fmt"

	"github.com/Aboagye-Dacosta/shopBackend/internal/database/models"
	"gorm.io/gorm"
)

func SeedAdminRole(db *gorm.DB) error {
	admin := models.Role{Name: "admin"}
	db.FirstOrCreate(&admin, models.Role{Name: "admin"})

	var perms models.Permission

	db.Where("name=?", "full_access").Find(&perms)
	if err := db.Model(&admin).Association("Permissions").Replace(perms); err != nil {
		return err
	}

	return nil
}

func SeedUserRole(db *gorm.DB) error {

	var userRole models.Role
	if err := db.FirstOrCreate(&userRole, models.Role{Name: "user"}).Error; err != nil {
		return err
	}

	permissions := []string{
		"view_products",
		"create_order",
		"view_orders",
		"create_payment",
		"cancel_order",
	}

	var existingPerms []models.Permission
	if err := db.Model(&userRole).Association("Permissions").Find(&existingPerms); err != nil {
		return err
	}

	existingMap := make(map[string]bool)
	for _, p := range existingPerms {
		existingMap[p.Name] = true
	}

	for _, name := range permissions {
		var perm models.Permission
		if err := db.FirstOrCreate(&perm, models.Permission{Name: name}).Error; err != nil {
			return err
		}

		if !existingMap[name] {
			if err := db.Model(&userRole).Association("Permissions").Append(&perm); err != nil {
				return err
			}
		}
	}

	fmt.Println("✅ User role seeded successfully with permissions")
	return nil
}
