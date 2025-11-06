package models

import (
	"github.com/lucsky/cuid"
	"gorm.io/gorm"
)

type Role struct {
	ID          string       `gorm:"type:char(25);primaryKey" json:"id"`
	Name        string       `gorm:"uniqueIndex;size:50;not null" json:"name" validate:"required,min=3,max=50"`
	Permissions []Permission `gorm:"many2many:role_permissions;" json:"permissions,omitempty"`
}

type Permission struct {
	ID   string `gorm:"type:char(25);primaryKey" json:"id"`
	Name string `gorm:"uniqueIndex;size:100;not null" json:"name" validate:"required,min=3,max=100"`
}

type PermissionsResponse struct {
	Response
	Data []Permission
}

type PermissionModel struct {
	DB *gorm.DB
}

func (r *Role) BeforeCreate(tx *gorm.DB) (err error) {
	if r.ID == "" {
		r.ID = cuid.New()
	}

	return
}

func (p *Permission) BeforeCreate(tx *gorm.DB) (err error) {
	if p.ID == "" {
		p.ID = cuid.New()
	}
	return
}
