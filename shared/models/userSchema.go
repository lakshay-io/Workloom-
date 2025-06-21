package models

import (
    "time"
    "gorm.io/gorm"
    "github.com/google/uuid"
)

const (
    ProviderGoogle = "google"
    ProviderGitHub = "github"
    ProviderLocal  = "local"
)

type User struct {
    ID           uuid.UUID      `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
    CreatedAt    time.Time
    UpdatedAt    time.Time
    DeletedAt    gorm.DeletedAt `gorm:"index"`

    Provider     string         `gorm:"type:varchar(20);not null"` 
    ProviderID   string         `gorm:"uniqueIndex:idx_provider_user;not null"` 
    Email        string         `gorm:"uniqueIndex;not null"`
    Name    string
    AvatarURL    string

    AccessToken  string         `gorm:"type:text"` 
    RefreshToken string         `gorm:"type:text"` 
    PasswordHash string         `gorm:"type:text"` 
    LastLogin    time.Time
}
