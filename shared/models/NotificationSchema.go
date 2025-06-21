package models

import (
	"time"

	"github.com/google/uuid"
)

type Notification struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
	UserID    uuid.UUID `gorm:"type:uuid;not null" json:"user_id"`

	// Type: ['mention', 'message', 'project_invite', 'task_assigned']
	Type      string    `gorm:"type:varchar(30);not null" json:"type"`
	EntityID  uuid.UUID `gorm:"type:uuid;not null" json:"entity_id"`
	Message   string    `gorm:"type:text" json:"message"`
	IsRead    bool      `gorm:"default:false" json:"is_read"`
	CreatedAt time.Time `json:"created_at"`

	// Associations
	User      User `gorm:"foreignKey:UserID" json:"-"`
}
