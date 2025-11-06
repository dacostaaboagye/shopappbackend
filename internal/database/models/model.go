package models

import "gorm.io/gorm"

type Models struct {
	Users    UserModel
	Products ProductModel
	Orders   OrderModel
	Payments PaymentModel
	Permissions PermissionModel
}

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Code    int         `json:"code"`
	Data    interface{} `json:"data,omitempty"`
}

type ErrResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Code    int         `json:"code"`
}

func NewModel(db *gorm.DB) *Models {
	return &Models{
		Users:    UserModel{db},
		Products: ProductModel{db},
		Payments: PaymentModel{db},
		Orders:   OrderModel{db},
		Permissions: PermissionModel{db},
	}
}
