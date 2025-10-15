package entity

import (
	"time"

	"github.com/google/uuid"
)

type RolePermission struct {
	RoleID       uuid.UUID `gorm:"type:uuid;primary_key;index:idx_role_permissions_role_id" json:"role_id"`
	PermissionID uuid.UUID `gorm:"type:uuid;primary_key;index:idx_role_permissions_permission_id" json:"permission_id"`
	CreatedAt    time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP" json:"created_at"`
}

func (RolePermission) TableName() string {
	return "role_permissions"
}
