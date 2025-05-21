package domain

import (
	"time"
	"github.com/google/uuid"
)

// Conversation represents a chat conversation with its details.
type Conversation struct {
	ID        uuid.UUID             `json:"id"`
	Type ConversationType `json:"type"`
	Name      *string               `json:"name,omitempty"`
	CreatedAt time.Time             `json:"created_at"`
	Members   []ConversationMember  `json:"members"`
	Messages  []Message             `json:"messages"`
}

// ConversationType defines the type of conversation, either private or channel.
type ConversationType string

// ConversationType constants
const (
	Private ConversationType = "PRIVATE"
	Channel ConversationType = "CHANNEL"
)