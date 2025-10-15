package repository

import (
	"context"

	"aolus-software/clean-gofiber/internal/entity"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepository interface {
	Repository
	Create(ctx context.Context, user *entity.User) error
	FindByID(ctx context.Context, id uuid.UUID) (*entity.User, error)
	FindByEmail(ctx context.Context, email string) (*entity.User, error)
	FindAll(ctx context.Context, limit, offset int) ([]entity.User, error)
	Update(ctx context.Context, user *entity.User) error
	Delete(ctx context.Context, id uuid.UUID) error
	Count(ctx context.Context) (int64, error)
	FindWithRoles(ctx context.Context, id uuid.UUID) (*entity.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) WithTx(tx *gorm.DB) Repository {
	return &userRepository{db: tx}
}

func (r *userRepository) getDB(ctx context.Context) *gorm.DB {
	if tx := GetTxFromContext(ctx); tx != nil {
		return tx
	}
	return r.db.WithContext(ctx)
}

func (r *userRepository) Create(ctx context.Context, user *entity.User) error {
	return r.getDB(ctx).Create(user).Error
}

func (r *userRepository) FindByID(ctx context.Context, id uuid.UUID) (*entity.User, error) {
	var user entity.User
	err := r.getDB(ctx).Where("id = ?", id).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) FindByEmail(ctx context.Context, email string) (*entity.User, error) {
	var user entity.User
	err := r.getDB(ctx).Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) FindAll(ctx context.Context, limit, offset int) ([]entity.User, error) {
	var users []entity.User
	err := r.getDB(ctx).Limit(limit).Offset(offset).Find(&users).Error
	return users, err
}

func (r *userRepository) Update(ctx context.Context, user *entity.User) error {
	return r.getDB(ctx).Save(user).Error
}

func (r *userRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.getDB(ctx).Delete(&entity.User{}, id).Error
}

func (r *userRepository) Count(ctx context.Context) (int64, error) {
	var count int64
	err := r.getDB(ctx).Model(&entity.User{}).Count(&count).Error
	return count, err
}

func (r *userRepository) FindWithRoles(ctx context.Context, id uuid.UUID) (*entity.User, error) {
	var user entity.User
	err := r.getDB(ctx).Preload("Roles").Where("id = ?", id).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
