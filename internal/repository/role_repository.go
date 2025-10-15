package repository

import (
	"context"

	"aolus-software/clean-gofiber/internal/entity"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type RoleRepository interface {
	Repository
	Create(ctx context.Context, role *entity.Role) error
	FindByID(ctx context.Context, id uuid.UUID) (*entity.Role, error)
	FindByName(ctx context.Context, name string) (*entity.Role, error)
	FindAll(ctx context.Context, limit, offset int) ([]entity.Role, error)
	Update(ctx context.Context, role *entity.Role) error
	Delete(ctx context.Context, id uuid.UUID) error
	Count(ctx context.Context) (int64, error)
	FindWithPermissions(ctx context.Context, id uuid.UUID) (*entity.Role, error)
	AttachPermissions(ctx context.Context, roleID uuid.UUID, permissionIDs []uuid.UUID) error
	DetachPermissions(ctx context.Context, roleID uuid.UUID, permissionIDs []uuid.UUID) error
	SyncPermissions(ctx context.Context, roleID uuid.UUID, permissionIDs []uuid.UUID) error
}

type roleRepository struct {
	db *gorm.DB
}

func NewRoleRepository(db *gorm.DB) RoleRepository {
	return &roleRepository{db: db}
}

func (r *roleRepository) WithTx(tx *gorm.DB) Repository {
	return &roleRepository{db: tx}
}

func (r *roleRepository) getDB(ctx context.Context) *gorm.DB {
	if tx := GetTxFromContext(ctx); tx != nil {
		return tx
	}
	return r.db.WithContext(ctx)
}

func (r *roleRepository) Create(ctx context.Context, role *entity.Role) error {
	return r.getDB(ctx).Create(role).Error
}

func (r *roleRepository) FindByID(ctx context.Context, id uuid.UUID) (*entity.Role, error) {
	var role entity.Role
	err := r.getDB(ctx).Where("id = ?", id).First(&role).Error
	if err != nil {
		return nil, err
	}
	return &role, nil
}

func (r *roleRepository) FindByName(ctx context.Context, name string) (*entity.Role, error) {
	var role entity.Role
	err := r.getDB(ctx).Where("name = ?", name).First(&role).Error
	if err != nil {
		return nil, err
	}
	return &role, nil
}

func (r *roleRepository) FindAll(ctx context.Context, limit, offset int) ([]entity.Role, error) {
	var roles []entity.Role
	err := r.getDB(ctx).Limit(limit).Offset(offset).Find(&roles).Error
	return roles, err
}

func (r *roleRepository) Update(ctx context.Context, role *entity.Role) error {
	return r.getDB(ctx).Save(role).Error
}

func (r *roleRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.getDB(ctx).Delete(&entity.Role{}, id).Error
}

func (r *roleRepository) Count(ctx context.Context) (int64, error) {
	var count int64
	err := r.getDB(ctx).Model(&entity.Role{}).Count(&count).Error
	return count, err
}

func (r *roleRepository) FindWithPermissions(ctx context.Context, id uuid.UUID) (*entity.Role, error) {
	var role entity.Role
	err := r.getDB(ctx).Preload("Permissions").Where("id = ?", id).First(&role).Error
	if err != nil {
		return nil, err
	}
	return &role, nil
}

func (r *roleRepository) AttachPermissions(ctx context.Context, roleID uuid.UUID, permissionIDs []uuid.UUID) error {
	db := r.getDB(ctx)
	for _, permID := range permissionIDs {
		rolePermission := entity.RolePermission{
			RoleID:       roleID,
			PermissionID: permID,
		}
		if err := db.Create(&rolePermission).Error; err != nil {
			return err
		}
	}
	return nil
}

func (r *roleRepository) DetachPermissions(ctx context.Context, roleID uuid.UUID, permissionIDs []uuid.UUID) error {
	return r.getDB(ctx).
		Where("role_id = ? AND permission_id IN ?", roleID, permissionIDs).
		Delete(&entity.RolePermission{}).Error
}

func (r *roleRepository) SyncPermissions(ctx context.Context, roleID uuid.UUID, permissionIDs []uuid.UUID) error {
	db := r.getDB(ctx)
	
	// Delete existing permissions
	if err := db.Where("role_id = ?", roleID).Delete(&entity.RolePermission{}).Error; err != nil {
		return err
	}
	
	// Add new permissions
	return r.AttachPermissions(ctx, roleID, permissionIDs)
}
