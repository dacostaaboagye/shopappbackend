package models

import (
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/lucsky/cuid"
	"gorm.io/gorm"
)

type PaymentModel struct {
	DB *gorm.DB
}

type Payment struct {
	ID          string     `json:"id" gorm:"primaryKey;size:36"`
	OrderID     string     `json:"order_id" gorm:"not null" validate:"required"`
	Amount      float64    `json:"amount" gorm:"not null" validate:"required,gt=0"`
	Method      string     `json:"method" gorm:"size:100;not null" validate:"required,oneof=card paypal mobile_money bank_transfer"`
	Status      string     `json:"status" gorm:"size:50;default:'pending'" validate:"required,oneof=pending completed failed refunded"`
	ProcessedAt *time.Time `json:"processed_at,omitempty"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`

	// Relations
	Order Order `json:"order"`
}

func (u *Payment) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = cuid.New()
	return
}

func (p *Payment) Validate() error {
	validate := validator.New()
	return validate.Struct(p)
}
