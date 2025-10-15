package entity

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID              uuid.UUID  `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	Name            string     `gorm:"type:varchar(255);not null" json:"name"`
	Email           string     `gorm:"type:varchar(255);not null;unique;index:idx_users_email" json:"email"`
	Password        string     `gorm:"type:varchar(255);not null" json:"-"`
	EmailVerifiedAt *time.Time `gorm:"type:timestamp" json:"email_verified_at"`
	CreatedAt       time.Time  `gorm:"type:timestamp;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt       time.Time  `gorm:"type:timestamp;default:CURRENT_TIMESTAMP" json:"updated_at"`
	Roles           []Role     `gorm:"many2many:user_roles;" json:"roles,omitempty"`
}

func (User) TableName() string {
	return "users"
}
