package seed

import (
	"github.com/Aboagye-Dacosta/shopBackend/internal/database/models"
	"gorm.io/gorm"
)

var DefaultPermissions = []string{
	//admin
	"full_access",
	// Products
	"view_products",
	"create_product",
	"update_product",
	"delete_product",
	"update_inventory",

	// Orders
	"view_orders",
	"create_order",
	"update_order_status",
	"cancel_order",
	"refund_order",

	// Payments
	"view_payments",
	"create_payment",
	"refund_payment",

	// Users
	"view_users",
	"create_user",
	"update_user",
	"delete_user",
	"ban_user",

	// Reports
	"view_reports",
	"export_data",

	// System
	"manage_roles",
	"manage_permissions",
	"manage_settings",
}

func SeedPermissions(db *gorm.DB) error {
	for _, name := range DefaultPermissions {
		var p models.Permission
		if err := db.Where(models.Permission{Name: name}).FirstOrCreate(&p, models.Permission{Name: name}).Error; err != nil {
			return err
		}
	}
	return nil
}
