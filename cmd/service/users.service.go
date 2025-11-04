package service

import (
	"context"
	"time"

	"github.com/Aboagye-Dacosta/shopBackend/internal/database/entities"
	"github.com/Aboagye-Dacosta/shopBackend/internal/database/models"
	"github.com/Aboagye-Dacosta/shopBackend/internal/errors"
	"github.com/Aboagye-Dacosta/shopBackend/logger"
	"github.com/Aboagye-Dacosta/shopBackend/utils"
)

const User = entities.USER

type UserService struct {
	Users *models.UserModel
}

func (s *UserService) GetAll(ctx context.Context) ([]*models.User, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	log := logger.FromContext(ctx)
	var users []*models.User
	if err := s.Users.DB.WithContext(ctx).Preload("Orders").Find(&users).Error; err != nil {
		log.ErrLogger.Error(err.Error(), "entity", User)
		return nil, errors.FromDb(User, err)
	}

	return users, nil
}

func (s *UserService) Create(ctx context.Context, user *models.RegisterRequest) (*models.User, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	log := logger.FromContext(ctx)
	var userValue models.User
	hashedPassword, err := utils.HashPassword(user.Password)

	if err != nil {
		log.ErrLogger.Error(err.Error(), "entity", User)
		return nil, errors.FromDb(User, err)
	}

	userValue.FirstName = user.FirstName
	userValue.LastName = user.LastName
	userValue.Email = user.Email
	userValue.Password = string(hashedPassword)

	if err := s.Users.DB.WithContext(ctx).Create(&userValue).Error; err != nil {
		log.ErrLogger.Error(err.Error(), "entity", User)
		return nil, errors.FromDb(User, err)
	}

	return &userValue, nil
}

func (s *UserService) GetByEmail(ctx context.Context, email string) (*models.User, error) {
	ctx, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	log := logger.FromContext(ctx)
	var user models.User
	if err :=
		s.Users.DB.WithContext(ctx).Where("email = ?", email).First(&user).Error; err != nil {
		log.ErrLogger.Error(err.Error(), "entity", User)
		return nil, errors.FromDb(User, err)
	}

	return &user, nil
}
