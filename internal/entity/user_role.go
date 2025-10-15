package entity

import (
	"time"

	"github.com/google/uuid"
)

type UserRole struct {
	UserID    uuid.UUID `gorm:"type:uuid;primary_key;index:idx_user_roles_user_id" json:"user_id"`
	RoleID    uuid.UUID `gorm:"type:uuid;primary_key;index:idx_user_roles_role_id" json:"role_id"`
	CreatedAt time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP" json:"created_at"`
}

func (UserRole) TableName() string {
	return "user_roles"
}
