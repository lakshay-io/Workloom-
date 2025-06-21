package models

import (
	"time"

	"github.com/google/uuid"
)

type Document struct {
	ID             uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
	UploadedByID   uuid.UUID `gorm:"type:uuid;not null" json:"uploaded_by_id"`
	ProjectID      uuid.UUID `gorm:"type:uuid;not null" json:"project_id"`
	URL            string    `gorm:"type:text;not null" json:"url"`
	Filename       string    `gorm:"type:text;not null" json:"filename"`
	FileType       string    `gorm:"type:text" json:"file_type"`
	SizeInBytes    int64     `json:"size_in_bytes"`
	CreatedAt      time.Time `json:"created_at"`

	// Associations
	UploadedBy     User    `gorm:"foreignKey:UploadedByID" json:"-"`
	Project        Project `gorm:"foreignKey:ProjectID" json:"-"`
}
