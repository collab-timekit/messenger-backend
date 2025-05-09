package ws

import (
    "time"
    "log"
)

// DispatchMessage routes an incoming message to the appropriate handler based on its type.
func DispatchMessage(sender *Client, msg IncomingMessage) {
    switch msg.Type {
		case MessageTypeChat:
			handleChatMessage(sender, msg)
		case MessageTypeStatus:
			handleStatusMessage(sender, msg)
		case MessageTypeTyping:
			handleTypingMessage(sender, msg)
		default:
            log.Println("Unknown message type:", msg.Type)
    }
}

func handleChatMessage(sender *Client, msg IncomingMessage) {
    out := OutgoingMessage{
        Type:     MessageTypeChat,
        Content:  msg.Content,
        FromUser: sender.userID,
        SentAt:   time.Now().Format(time.RFC3339),
    }

    if receiver, ok := sender.hub.GetClient(msg.ToUser); ok {
        receiver.send <- out
    }
}

func handleStatusMessage(sender *Client, msg IncomingMessage) {
    // Np. update "online" status i broadcast do kontaktów
}

func handleTypingMessage(sender *Client, msg IncomingMessage) {
    if receiver, ok := sender.hub.GetClient(msg.ToUser); ok {
        receiver.send <- OutgoingMessage{
            Type:     MessageTypeTyping,
            Content:  "typing...",
            FromUser: sender.userID,
            SentAt:   time.Now().Format(time.RFC3339),
        }
    }
}