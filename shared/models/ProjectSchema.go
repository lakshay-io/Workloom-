package models

import (
	"time"

	"github.com/google/uuid"
)

type Project struct {
	ID           uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
	WorkspaceID  uuid.UUID `gorm:"type:uuid;not null" json:"workspace_id"`
	Name         string    `gorm:"type:text;not null" json:"name"`
	Description  string    `gorm:"type:text" json:"description"`
	CreatedByID  uuid.UUID `gorm:"type:uuid;not null" json:"created_by_id"`

	// Status: 'active', 'archived', 'completed'
	Status       string     `gorm:"type:varchar(20);default:'active'" json:"status"`
	Deadline     *time.Time `json:"deadline"`

	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`

	// Associations
	CreatedBy    User      `gorm:"foreignKey:CreatedByID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"-"`
	Workspace    Workspace `gorm:"foreignKey:WorkspaceID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"-"`
}

type ProjectMember struct {
	ID         uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
	ProjectID  uuid.UUID `gorm:"type:uuid;not null" json:"project_id"`
	UserID     uuid.UUID `gorm:"type:uuid;not null" json:"user_id"`

	// Role: 'owner', 'editor', 'viewer'
	Role       string    `gorm:"type:varchar(20);not null" json:"role"`

	AddedByID  uuid.UUID `gorm:"type:uuid;not null" json:"added_by_id"`
	JoinedAt   time.Time `json:"joined_at"`

	// Associations
	Project    Project `gorm:"foreignKey:ProjectID" json:"-"`
	User       User    `gorm:"foreignKey:UserID" json:"-"`
	AddedBy    User    `gorm:"foreignKey:AddedByID" json:"-"`
}
