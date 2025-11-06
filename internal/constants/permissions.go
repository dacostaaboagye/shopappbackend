package constants

type Permission string

const (
	// Admin
	FullAccess Permission = "full_access"

	// Products
	ViewProducts     Permission = "view_products"
	CreateProduct    Permission = "create_product"
	UpdateProduct    Permission = "update_product"
	DeleteProduct    Permission = "delete_product"
	UpdateInventory  Permission = "update_inventory"

	// Orders
	ViewOrders        Permission = "view_orders"
	CreateOrder       Permission = "create_order"
	UpdateOrderStatus Permission = "update_order_status"
	CancelOrder       Permission = "cancel_order"
	RefundOrder       Permission = "refund_order"

	// Payments
	ViewPayments   Permission = "view_payments"
	CreatePayment  Permission = "create_payment"
	RefundPayment  Permission = "refund_payment"

	// Users
	ViewUsers   Permission = "view_users"
	CreateUser  Permission = "create_user"
	UpdateUser  Permission = "update_user"
	DeleteUser  Permission = "delete_user"
	BanUser     Permission = "ban_user"

	// Reports
	ViewReports Permission = "view_reports"
	ExportData  Permission = "export_data"

	// System
	ManageRoles       Permission = "manage_roles"
	ManagePermissions Permission = "manage_permissions"
	ManageSettings    Permission = "manage_settings"
)
