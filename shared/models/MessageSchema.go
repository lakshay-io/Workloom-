package models

import (
	"time"

	"github.com/google/uuid"
)

type Message struct {
	ID              uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
	ConversationID  uuid.UUID `gorm:"type:uuid;not null" json:"conversation_id"`
	SenderID        uuid.UUID `gorm:"type:uuid;not null" json:"sender_id"`
	Content         string    `gorm:"type:text" json:"content"`

	// MessageType: ['text', 'image', 'file', 'code', 'thread', 'system']
	MessageType     string     `gorm:"type:varchar(20);not null" json:"message_type"`
	ThreadParentID  *uuid.UUID `gorm:"type:uuid" json:"thread_parent_id"`
	AttachmentURL   string     `gorm:"type:text" json:"attachment_url"`
	CreatedAt       time.Time  `json:"created_at"`
	UpdatedAt       time.Time  `json:"updated_at"`

	// Associations
	Conversation    Conversation `gorm:"foreignKey:ConversationID" json:"-"`
	Sender          User         `gorm:"foreignKey:SenderID" json:"-"`
	Parent          *Message     `gorm:"foreignKey:ThreadParentID" json:"-"`
}

type MessageReaction struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
	MessageID uuid.UUID `gorm:"type:uuid;not null" json:"message_id"`
	UserID    uuid.UUID `gorm:"type:uuid;not null" json:"user_id"`
	Reaction  string    `gorm:"type:varchar(10);not null" json:"reaction"`
	CreatedAt time.Time `json:"created_at"`

	// Associations
	Message   Message `gorm:"foreignKey:MessageID" json:"-"`
	User      User    `gorm:"foreignKey:UserID" json:"-"`
}
