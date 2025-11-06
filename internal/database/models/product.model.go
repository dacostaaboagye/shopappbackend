package models

import (
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/lucsky/cuid"
	"gorm.io/gorm"
)

type ProductModel struct {
	DB *gorm.DB
}

type Product struct {
	ID          string    `json:"id" gorm:"primaryKey;size:36"`
	Name        string    `json:"name" gorm:"size:255;not null" validate:"required,min=2"`
	Description string    `json:"description" gorm:"type:text" validate:"omitempty"`
	Price       float64   `json:"price" gorm:"not null" validate:"required,gt=0"`
	Stock       int       `json:"stock" gorm:"not null" validate:"gte=0"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`

	// Relations
	Orders []Order `json:"orders,omitempty" gorm:"many2many:order_products;"`
}

func (p *Product) BeforeCreate(tx *gorm.DB) (err error) {
	if p.ID == "" {
		p.ID = cuid.New()
	}
	return
}

func (p *Product) Validate() error {
	validate := validator.New()
	return validate.Struct(p)
}
