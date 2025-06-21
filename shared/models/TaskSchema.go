package models

import (
	"time"

	"github.com/google/uuid"
)

type Task struct {
	ID            uuid.UUID  `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
	ProjectID     uuid.UUID  `gorm:"type:uuid;not null" json:"project_id"`
	AssignedToID  uuid.UUID  `gorm:"type:uuid" json:"assigned_to_id"`
	Title         string     `gorm:"type:text;not null" json:"title"`
	Description   string     `gorm:"type:text" json:"description"`
	Status        string     `gorm:"type:varchar(20);default:'todo'" json:"status"`       // ['todo', 'in_progress', 'done']
	Priority      string     `gorm:"type:varchar(10);default:'medium'" json:"priority"`   // ['low', 'medium', 'high']
	DueDate       *time.Time `json:"due_date"`
	CreatedByID   uuid.UUID  `gorm:"type:uuid;not null" json:"created_by_id"`
	CreatedAt     time.Time  `json:"created_at"`
	UpdatedAt     time.Time  `json:"updated_at"`

	// Associations (optional)
	Project     Project `gorm:"foreignKey:ProjectID" json:"-"`
	AssignedTo  User    `gorm:"foreignKey:AssignedToID" json:"-"`
	CreatedBy   User    `gorm:"foreignKey:CreatedByID" json:"-"`
}


type TaskComment struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
	TaskID    uuid.UUID `gorm:"type:uuid;not null" json:"task_id"`
	UserID    uuid.UUID `gorm:"type:uuid;not null" json:"user_id"`
	Content   string    `gorm:"type:text;not null" json:"content"`
	CreatedAt time.Time `json:"created_at"`

	// Associations
	Task      Task `gorm:"foreignKey:TaskID" json:"-"`
	User      User `gorm:"foreignKey:UserID" json:"-"`
}
