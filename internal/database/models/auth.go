package models

import (
	"github.com/go-playground/validator/v10"
)

type RegisterRequest struct {
	FirstName string `json:"first_name" validate:"required,min=2"`
	LastName  string `json:"last_name" validate:"required,min=2"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,min=8"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

type AuthResponse struct {
	Token string `json:"token"`
	User  User   `json:"user"`
}

func (r *RegisterRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(r)
}

func (r *LoginRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(r)
}
