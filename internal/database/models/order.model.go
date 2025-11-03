package models

import (
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/lucsky/cuid"
	"gorm.io/gorm"
)

type OrderModel struct {
	DB *gorm.DB
}

type Order struct {
	ID         string    `json:"id" gorm:"primaryKey;size:36"`
	UserID     string    `json:"user_id" validate:"required"`
	Status     string    `json:"status" gorm:"size:50;default:'pending'" validate:"required,oneof=pending paid shipped delivered canceled"`
	TotalPrice float64   `json:"total_price" validate:"required,gt=0"`
	OrderedAt  time.Time `json:"ordered_at"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`

	// Relations
	User     User      `json:"user"`
	Products []Product `json:"products" gorm:"many2many:order_products;"`
	Payments []Payment `json:"payments,omitempty" gorm:"foreignKey:OrderID"`
}

func (u *Order) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = cuid.New()
	return
}

func (o *Order) Validate() error {
	validate := validator.New()
	return validate.Struct(o)
}
