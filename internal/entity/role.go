package entity

import (
	"time"

	"github.com/google/uuid"
)

type Role struct {
	ID          uuid.UUID    `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	Name        string       `gorm:"type:varchar(255);not null;unique;index:idx_roles_name" json:"name"`
	CreatedAt   time.Time    `gorm:"type:timestamp;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time    `gorm:"type:timestamp;default:CURRENT_TIMESTAMP" json:"updated_at"`
	Users       []User       `gorm:"many2many:user_roles;" json:"users,omitempty"`
	Permissions []Permission `gorm:"many2many:role_permissions;" json:"permissions,omitempty"`
}

func (Role) TableName() string {
	return "roles"
}
