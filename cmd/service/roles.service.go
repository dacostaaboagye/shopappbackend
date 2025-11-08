package service

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/Aboagye-Dacosta/shopBackend/internal/codes"
	"github.com/Aboagye-Dacosta/shopBackend/internal/database/entities"
	"github.com/Aboagye-Dacosta/shopBackend/internal/database/models"
	appErrors "github.com/Aboagye-Dacosta/shopBackend/internal/errors"
	"github.com/Aboagye-Dacosta/shopBackend/internal/logger"
	"gorm.io/gorm"
)

const Role = entities.ROLE

type RoleService struct {
	roles *models.RoleModel
}

// Get All Roles
func (s *RoleService) GetAll(ctx context.Context) (roles []*models.Role, err error) {
	ctx, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	log := logger.FromContext(ctx)
	if err = s.roles.DB.WithContext(ctx).Preload("Permissions").Find(roles).Error; err != nil {
		log.ErrLogger.ErrorContext(ctx, err.Error(), "entity", User)
		err = appErrors.FromDb(Role, err)
	}

	return

}

// Get Single Role
func (s *RoleService) GetRole(ctx context.Context, id string) (tRole *models.Role, err error) {
	ctx, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	log := logger.FromContext(ctx)
	if err = s.roles.DB.WithContext(ctx).Where("id=?", id).Preload("Permissions").First(tRole).Error; err != nil {
		log.ErrLogger.ErrorContext(ctx, err.Error(), "entity", User)
		err = appErrors.FromDb(Role, err)
	}

	return
}

// Create Role
func (s *RoleService) CreateRole(ctx context.Context, role *models.RoleRequest) (*models.Role, error) {
	ctx, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()
	log := logger.FromContext(ctx)

	var existing models.Role
	err := s.roles.DB.WithContext(ctx).Where("name = ?", role.Role).First(&existing).Error

	if err == nil {
		return nil, appErrors.New(Role, http.StatusConflict, errors.New("Role already exist"))
	}

	if !errors.Is(err, gorm.ErrRecordNotFound) {
		log.ErrLogger.ErrorContext(ctx, err.Error(), "entity", Role)
		return nil, appErrors.FromDb(Role, err)
	}

	var roleResult models.Role
	// Transaction for role creation + permission assignment
	err = s.roles.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		newRole := models.Role{Name: role.Role}
		if err := tx.Create(&newRole).Error; err != nil {
			return err
		}

		var permissions []models.Permission
		if err := tx.Where("name IN ?", role.Permissions).Find(&permissions).Error; err != nil {
			return err
		}

		if len(permissions) > 0 {
			if err := tx.Model(&newRole).Association("Permissions").Append(permissions); err != nil {
				return err
			}
			newRole.Permissions = permissions
		}

		// Return via closure capture
		roleResult = newRole
		return nil
	})

	if err != nil {
		log.ErrLogger.ErrorContext(ctx, err.Error(), "entity", Role)
		return nil, appErrors.FromDb(Role, err)
	}

	return &roleResult, nil
}

// Update Role
func (s *RoleService) UpdateRole(ctx context.Context, id string, role *models.RoleRequest) (*models.Role, error) {
	ctx, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	//check exist
	var existing models.Role
	err := s.roles.DB.WithContext(ctx).Where("id=?", id).First(&existing).Error

	if err != nil {
		return nil, appErrors.FromDb(Role, err)
	}

	//update the role name
	err = s.roles.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err = tx.Model(&existing).Association("Permissions").Clear(); err != nil {
			return err
		}

		if err = tx.Model(&existing).Update("name", role.Role).Error; err != nil {
			return err
		}

		var permissions []models.Permission
		if err := tx.Where("name IN ?", role.Permissions).Find(&permissions).Error; err != nil {
			return err
		}

		if err := tx.Model(&existing).Association("Permissions").Append(permissions); err != nil {
			return err
		}

		existing.Permissions = permissions

		return nil
	})

	if err != nil {
		return nil, appErrors.FromDb(Role, err)
	}

	// Reload updated record with associations
	if err := s.roles.DB.WithContext(ctx).
		Preload("Permissions").
		First(&existing, "id = ?", id).Error; err != nil {
		return nil, appErrors.FromDb(Role, err)
	}

	return &existing, nil

}

// Delete Role
func (s *RoleService) DeleteRole(ctx context.Context, id string) error {
	ctx, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	// Step 1: Check if the role exists
	var existing models.Role
	if err := s.roles.DB.WithContext(ctx).Where("id = ?", id).First(&existing).Error; err != nil {
		return appErrors.FromDb(Role, err)
	}

	// Step 2: Check if any user has this role
	var count int64
	if err := s.roles.DB.WithContext(ctx).
		Model(&models.User{}).
		Joins("JOIN user_roles ur ON ur.user_id = users.id").
		Where("ur.role_id = ?", id).
		Count(&count).Error; err != nil {
		return appErrors.FromDb(Role, err)
	}

	if count > 0 {
		return appErrors.New(Role, codes.ROLE_IN_USE, errors.New("Roles is in use"))
	}

	err := s.roles.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// Clear role-permission relationships
		if err := tx.Model(&existing).Association("Permissions").Clear(); err != nil {
			return err
		}

		// Delete the role
		if err := tx.Delete(&existing).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return appErrors.FromDb(Role, err)
	}

	return nil
}
