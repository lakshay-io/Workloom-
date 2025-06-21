package models

import (
	"time"

	"github.com/google/uuid"
)

type Conversation struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
	ProjectID   uuid.UUID `gorm:"type:uuid;not null" json:"project_id"`
	IsGroup     bool      `gorm:"not null" json:"is_group"`
	Title       string    `gorm:"type:text" json:"title"`
	CreatedByID uuid.UUID `gorm:"type:uuid;not null" json:"created_by_id"`
	CreatedAt   time.Time `json:"created_at"`

	// Associations
	Project     Project `gorm:"foreignKey:ProjectID" json:"-"`
	CreatedBy   User    `gorm:"foreignKey:CreatedByID" json:"-"`
}

type ConversationParticipant struct {
	ID             uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
	ConversationID uuid.UUID `gorm:"type:uuid;not null" json:"conversation_id"`
	UserID         uuid.UUID `gorm:"type:uuid;not null" json:"user_id"`
	JoinedAt       time.Time `json:"joined_at"`
	LastSeenAt     time.Time `json:"last_seen_at"`

	// Associations
	Conversation   Conversation `gorm:"foreignKey:ConversationID" json:"-"`
	User           User         `gorm:"foreignKey:UserID" json:"-"`
}
