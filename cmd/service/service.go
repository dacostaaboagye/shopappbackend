package service

import "github.com/Aboagye-Dacosta/shopBackend/internal/database/models"

type Service struct {
	UserService       *UserService
	PermissionService *PermissionService
}

func NewService(m *models.Models) *Service {
	return &Service{
		UserService:       &UserService{&m.Users},
		PermissionService: &PermissionService{&m.Permissions},
	}
}
