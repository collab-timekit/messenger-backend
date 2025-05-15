package domain

import (
	"time"

	"github.com/google/uuid"
)

type Conversation struct {
	ID        uuid.UUID
	Type      string
	Name      *string
	CreatedAt time.Time
	Members   []ConversationMember
	Messages  []Message
}