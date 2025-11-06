package service

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/Aboagye-Dacosta/shopBackend/internal/database/entities"
	"github.com/Aboagye-Dacosta/shopBackend/internal/database/models"
	appErrors "github.com/Aboagye-Dacosta/shopBackend/internal/errors"
	"github.com/Aboagye-Dacosta/shopBackend/internal/logger"
	"github.com/Aboagye-Dacosta/shopBackend/internal/utils"
	"gorm.io/gorm"
)

const User = entities.USER

type UserService struct {
	*models.UserModel
}

func (s *UserService) GetAll(ctx context.Context) ([]*models.User, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	log := logger.FromContext(ctx)
	var users []*models.User
	if err := s.DB.WithContext(ctx).Preload("Roles").Preload("Orders").Find(&users).Error; err != nil {
		log.ErrLogger.Error(err.Error(), "entity", User)
		return nil, appErrors.FromDb(User, err)
	}

	return users, nil
}

func (s *UserService) GetById(ctx context.Context, id string) (*models.User, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	log := logger.FromContext(ctx)
	var user models.User
	if err := s.DB.WithContext(ctx).Where("id=?", id).
		Preload("Roles.Permissions").
		Preload("Roles").
		First(&user).Error; err != nil {
		log.ErrLogger.Error(err.Error(), "entity", User)
		return nil, appErrors.FromDb(User, err)
	}

	return &user, nil
}

func (s *UserService) Create(ctx context.Context, req *models.RegisterRequest) (*models.User, error) {
	ctx, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	log := logger.FromContext(ctx)

	var existing models.User
	if err := s.DB.WithContext(ctx).
		Select("id").
		Where("email = ?", req.Email).
		First(&existing).Error; err == nil {
		return nil, appErrors.New(User, http.StatusConflict, fmt.Errorf("user with email %s already exists", req.Email))
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, appErrors.FromDb(User, err)
	}

	hashed, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, appErrors.FromDb(User, err)
	}

	user := &models.User{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
		Password:  hashed,
	}

	if err := s.DB.WithContext(ctx).Create(&user).Error; err != nil {
		return nil, appErrors.FromDb(User, err)
	}

	var role models.Role
	if err := s.DB.WithContext(ctx).
		Preload("Permissions").
		Where("name = ?", "user").
		First(&role).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, appErrors.New(User, http.StatusInternalServerError, fmt.Errorf("default 'user' role not found"))
		}
		return nil, appErrors.FromDb(User, err)
	}

	if err := s.DB.WithContext(ctx).Model(&user).Association("Roles").Append(&role); err != nil {
		return nil, appErrors.FromDb(User, err)
	}

	if err := s.DB.WithContext(ctx).
		Preload("Roles.Permissions").
		First(&user, "id = ?", user.ID).Error; err != nil {
		return nil, appErrors.FromDb(User, err)
	}

	log.InfoLogger.Info("User created successfully", "userID", user.ID)
	return user, nil
}

func (s *UserService) GetByEmail(ctx context.Context, email string) (*models.User, error) {
	ctx, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	log := logger.FromContext(ctx)
	var user models.User
	if err := s.DB.WithContext(ctx).
		Where("email = ?", email).
		Preload("Roles.Permissions").
		Preload("Roles").
		First(&user).Error; err != nil {

		log.ErrLogger.Error(err.Error(), "entity", User)
		return nil, appErrors.FromDb(User, err)
	}

	return &user, nil
}
