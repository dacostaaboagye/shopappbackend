package controller

import (
	"github.com/Aboagye-Dacosta/shopBackend/cmd/service"
)

type Controller struct {
	userService *service.UserService
	permissionService *service.PermissionService
}

func NewController(s *service.Service) *Controller {
	return &Controller{userService: s.UserService,permissionService: s.PermissionService}
}
