package repository

import (
	"context"

	"aolus-software/clean-gofiber/internal/entity"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PermissionRepository interface {
	Repository
	Create(ctx context.Context, permission *entity.Permission) error
	FindByID(ctx context.Context, id uuid.UUID) (*entity.Permission, error)
	FindByName(ctx context.Context, name string) (*entity.Permission, error)
	FindAll(ctx context.Context, limit, offset int) ([]entity.Permission, error)
	FindByGroup(ctx context.Context, group string) ([]entity.Permission, error)
	Update(ctx context.Context, permission *entity.Permission) error
	Delete(ctx context.Context, id uuid.UUID) error
	Count(ctx context.Context) (int64, error)
	FindByIDs(ctx context.Context, ids []uuid.UUID) ([]entity.Permission, error)
}

type permissionRepository struct {
	db *gorm.DB
}

func NewPermissionRepository(db *gorm.DB) PermissionRepository {
	return &permissionRepository{db: db}
}

func (r *permissionRepository) WithTx(tx *gorm.DB) Repository {
	return &permissionRepository{db: tx}
}

func (r *permissionRepository) getDB(ctx context.Context) *gorm.DB {
	if tx := GetTxFromContext(ctx); tx != nil {
		return tx
	}
	return r.db.WithContext(ctx)
}

func (r *permissionRepository) Create(ctx context.Context, permission *entity.Permission) error {
	return r.getDB(ctx).Create(permission).Error
}

func (r *permissionRepository) FindByID(ctx context.Context, id uuid.UUID) (*entity.Permission, error) {
	var permission entity.Permission
	err := r.getDB(ctx).Where("id = ?", id).First(&permission).Error
	if err != nil {
		return nil, err
	}
	return &permission, nil
}

func (r *permissionRepository) FindByName(ctx context.Context, name string) (*entity.Permission, error) {
	var permission entity.Permission
	err := r.getDB(ctx).Where("name = ?", name).First(&permission).Error
	if err != nil {
		return nil, err
	}
	return &permission, nil
}

func (r *permissionRepository) FindAll(ctx context.Context, limit, offset int) ([]entity.Permission, error) {
	var permissions []entity.Permission
	err := r.getDB(ctx).Limit(limit).Offset(offset).Find(&permissions).Error
	return permissions, err
}

func (r *permissionRepository) FindByGroup(ctx context.Context, group string) ([]entity.Permission, error) {
	var permissions []entity.Permission
	err := r.getDB(ctx).Where("\"group\" = ?", group).Find(&permissions).Error
	return permissions, err
}

func (r *permissionRepository) Update(ctx context.Context, permission *entity.Permission) error {
	return r.getDB(ctx).Save(permission).Error
}

func (r *permissionRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.getDB(ctx).Delete(&entity.Permission{}, id).Error
}

func (r *permissionRepository) Count(ctx context.Context) (int64, error) {
	var count int64
	err := r.getDB(ctx).Model(&entity.Permission{}).Count(&count).Error
	return count, err
}

func (r *permissionRepository) FindByIDs(ctx context.Context, ids []uuid.UUID) ([]entity.Permission, error) {
	var permissions []entity.Permission
	err := r.getDB(ctx).Where("id IN ?", ids).Find(&permissions).Error
	return permissions, err
}
