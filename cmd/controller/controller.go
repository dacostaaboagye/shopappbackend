package controller

import (
	"github.com/Aboagye-Dacosta/shopBackend/cmd/service"
)

type Controller struct {
	userService *service.UserService
}

func NewController(s *service.Service) *Controller {
	return &Controller{userService: s.UserService}
}
