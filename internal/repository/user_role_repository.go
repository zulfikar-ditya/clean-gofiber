package repository

import (
	"context"

	"aolus-software/clean-gofiber/internal/entity"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRoleRepository interface {
	Repository
	AttachRoles(ctx context.Context, userID uuid.UUID, roleIDs []uuid.UUID) error
	DetachRoles(ctx context.Context, userID uuid.UUID, roleIDs []uuid.UUID) error
	SyncRoles(ctx context.Context, userID uuid.UUID, roleIDs []uuid.UUID) error
	GetUserRoles(ctx context.Context, userID uuid.UUID) ([]entity.Role, error)
	GetRoleUsers(ctx context.Context, roleID uuid.UUID) ([]entity.User, error)
	HasRole(ctx context.Context, userID uuid.UUID, roleID uuid.UUID) (bool, error)
}

type userRoleRepository struct {
	db *gorm.DB
}

func NewUserRoleRepository(db *gorm.DB) UserRoleRepository {
	return &userRoleRepository{db: db}
}

func (r *userRoleRepository) WithTx(tx *gorm.DB) Repository {
	return &userRoleRepository{db: tx}
}

func (r *userRoleRepository) getDB(ctx context.Context) *gorm.DB {
	if tx := GetTxFromContext(ctx); tx != nil {
		return tx
	}
	return r.db.WithContext(ctx)
}

func (r *userRoleRepository) AttachRoles(ctx context.Context, userID uuid.UUID, roleIDs []uuid.UUID) error {
	db := r.getDB(ctx)
	for _, roleID := range roleIDs {
		userRole := entity.UserRole{
			UserID: userID,
			RoleID: roleID,
		}
		if err := db.Create(&userRole).Error; err != nil {
			return err
		}
	}
	return nil
}

func (r *userRoleRepository) DetachRoles(ctx context.Context, userID uuid.UUID, roleIDs []uuid.UUID) error {
	return r.getDB(ctx).
		Where("user_id = ? AND role_id IN ?", userID, roleIDs).
		Delete(&entity.UserRole{}).Error
}

func (r *userRoleRepository) SyncRoles(ctx context.Context, userID uuid.UUID, roleIDs []uuid.UUID) error {
	db := r.getDB(ctx)
	
	// Delete existing roles
	if err := db.Where("user_id = ?", userID).Delete(&entity.UserRole{}).Error; err != nil {
		return err
	}
	
	// Add new roles
	return r.AttachRoles(ctx, userID, roleIDs)
}

func (r *userRoleRepository) GetUserRoles(ctx context.Context, userID uuid.UUID) ([]entity.Role, error) {
	var roles []entity.Role
	err := r.getDB(ctx).
		Joins("JOIN user_roles ON user_roles.role_id = roles.id").
		Where("user_roles.user_id = ?", userID).
		Find(&roles).Error
	return roles, err
}

func (r *userRoleRepository) GetRoleUsers(ctx context.Context, roleID uuid.UUID) ([]entity.User, error) {
	var users []entity.User
	err := r.getDB(ctx).
		Joins("JOIN user_roles ON user_roles.user_id = users.id").
		Where("user_roles.role_id = ?", roleID).
		Find(&users).Error
	return users, err
}

func (r *userRoleRepository) HasRole(ctx context.Context, userID uuid.UUID, roleID uuid.UUID) (bool, error) {
	var count int64
	err := r.getDB(ctx).
		Model(&entity.UserRole{}).
		Where("user_id = ? AND role_id = ?", userID, roleID).
		Count(&count).Error
	return count > 0, err
}
