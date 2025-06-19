package models

import (
    "gorm.io/gorm"
    "time"
)

type AuthProvider string

const (
    Google AuthProvider = "google"
    GitHub AuthProvider = "github"
)

type User struct {
    gorm.Model
    Provider     AuthProvider `gorm:"type:varchar(20);not null"` 
    ProviderID   string       `gorm:"uniqueIndex;not null"`       
    Name         string
    Email        string       `gorm:"uniqueIndex;not null"`
    AvatarURL    string
    AccessToken  string       `gorm:"type:text"` 
    RefreshToken string       `gorm:"type:text"` 
    LastLogin    time.Time
}
