package models

import (
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/lucsky/cuid"
	"gorm.io/gorm"
)

type Customer struct {
	ID           string    `json:"id" gorm:"primaryKey;size:36"`
	UserID       *string   `json:"user_id,omitempty" gorm:"uniqueIndex" validate:"omitempty,uuid4"`
	CustomerCard string    `json:"customer_card" gorm:"uniqueIndex;not null" validate:"required,min=6"`
	Points       int       `json:"points" gorm:"default:0" validate:"gte=0"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

func (c *Customer) BeforeCreate(tx *gorm.DB) (err error) {
	if c.ID == "" {
		c.ID = cuid.New()
	}
	return
}

func (c *Customer) Validate() error {
	validate := validator.New()
	return validate.Struct(c)
}
