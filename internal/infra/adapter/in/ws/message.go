package ws

// MessageTypeChat constants
const (
    MessageTypeChat   string = "chat"
    MessageTypeStatus string = "status"
    MessageTypeTyping string = "typing"
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