package entity

import (
	"time"

	"github.com/google/uuid"
)

type Permission struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	Name      string    `gorm:"type:varchar(255);not null;index:idx_permissions_name" json:"name"`
	Group     string    `gorm:"type:varchar(255);not null;index:idx_permissions_group;column:group" json:"group"`
	CreatedAt time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP" json:"updated_at"`
	Roles     []Role    `gorm:"many2many:role_permissions;" json:"roles,omitempty"`
}

func (Permission) TableName() string {
	return "permissions"
}
