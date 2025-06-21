package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/datatypes"
)

type AuditLog struct {
	ID           uuid.UUID      `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
	UserID       uuid.UUID      `gorm:"type:uuid;not null" json:"user_id"`
	WorkspaceID  uuid.UUID      `gorm:"type:uuid;not null" json:"workspace_id"`
	ProjectID    uuid.UUID      `gorm:"type:uuid" json:"project_id"`
	ActionType   string         `gorm:"type:text;not null" json:"action_type"`     // ['created', 'updated', 'deleted', etc.]
	EntityType   string         `gorm:"type:text;not null" json:"entity_type"`     // ['task', 'project', 'message', etc.]
	EntityID     uuid.UUID      `gorm:"type:uuid;not null" json:"entity_id"`
	Metadata     datatypes.JSON `gorm:"type:jsonb" json:"metadata"`
	CreatedAt    time.Time      `json:"created_at"`

	// Associations
	User         User      `gorm:"foreignKey:UserID" json:"-"`
	Workspace    Workspace `gorm:"foreignKey:WorkspaceID" json:"-"`
	Project      Project   `gorm:"foreignKey:ProjectID" json:"-"`
}
