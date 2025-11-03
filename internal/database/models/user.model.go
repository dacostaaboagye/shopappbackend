package models

import (
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/lucsky/cuid"
	"gorm.io/gorm"
)

type UserModel struct {
	DB *gorm.DB
}

type User struct {
	ID          string     `json:"id" gorm:"primaryKey;size:36"`
	FirstName   string     `json:"first_name" gorm:"size:100;not null" validate:"required,min=2"`
	LastName    string     `json:"last_name" gorm:"size:100;not null" validate:"required,min=2"`
	Email       string     `json:"email" gorm:"uniqueIndex;not null" validate:"required,email"`
	Password    string     `json:"-" gorm:"not null" validate:"required,min=8"`
	Birthday    *time.Time `json:"birthday,omitempty" validate:"omitempty,lte"`
	ActivatedAt *time.Time `json:"activated_at,omitempty"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`

	Orders []Order `json:"orders,omitempty" gorm:"foreignKey:UserID"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = cuid.New()
	return
}

func (u *User) Validate() error {
	validate := validator.New()
	return validate.Struct(u)
}
