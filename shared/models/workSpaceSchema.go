package models

import (
	"time"
    "github.com/google/uuid"
)

type Workspace struct {
	ID          uuid.UUID      `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
	Title       string         `gorm:"type:varchar(200);not null" json:"title"`
	Description string         `gorm:"type:text" json:"description"`
	CreatedByID uint           `gorm:"not null" json:"created_by_id"`
	LogoURL     string         `gorm:"type:text" json:"logo_url"`
	InviteCode  string         `gorm:"uniqueIndex" json:"invite_code"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`

	// Associations
	CreatedBy   User             `gorm:"foreignKey:CreatedByID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"-"`
}

type WorkspaceMember struct {
	ID          uuid.UUID    `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
	WorkspaceID uuid.UUID    `gorm:"type:uuid;not null" json:"workspace_id"`
	UserID      uuid.UUID    `gorm:"type:uuid;not null" json:"user_id"`

	// Roles: 'admin', 'member', 'viewer'
	Role        string       `gorm:"type:varchar(20);not null" json:"role"`

	// Status: 'active', 'invited', 'removed'
	Status      string       `gorm:"type:varchar(20);not null;default:'invited'" json:"status"`
	JoinedAt    time.Time    `json:"joined_at"`

	// Associations
	User        User        `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"-"`
	Workspace   Workspace   `gorm:"foreignKey:WorkspaceID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"-"`
}
