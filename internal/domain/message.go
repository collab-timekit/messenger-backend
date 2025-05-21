package domain

import (
    "time"
    "github.com/google/uuid"
)

// Message represents a message in a conversation.
type Message struct {
    ID             uuid.UUID `json:"id"`
    ConversationID uuid.UUID `json:"conversation_id"`
    SenderID       uuid.UUID `json:"sender_id"`
    Content        string    `json:"content"`
    CreatedAt      time.Time `json:"created_at"`
    Edited         bool      `json:"edited"`
}