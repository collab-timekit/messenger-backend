package ws

const (
    // MessageTypeChat represents a chat message type.
    MessageTypeChat   string = "chat"
	// MessageTypeStatus represents a status message type.
    MessageTypeStatus string = "status"
	// MessageTypeTyping represents a typing message type.
    MessageTypeTyping string = "typing"
	// MessageTypeError represents an error message type.
    MessageTypeError  string = "error"
)

// IncomingMessage represents a message received via WebSocket.
type IncomingMessage struct {
    Type    string         `json:"type"`
    Content string         `json:"content,omitempty"`
    ToUser  string         `json:"toUser,omitempty"`
}

// OutgoingMessage represents a message sent via WebSocket.
type OutgoingMessage struct {
    Type     string         `json:"type"`
    Content  string         `json:"content"`
    FromUser string         `json:"fromUser"`
    SentAt   string         `json:"sentAt"`
}