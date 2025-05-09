package ws

import (
    "github.com/gorilla/websocket"
    "log"
    "time"
)

// Client represents a WebSocket client connected to the hub.
type Client struct {
    userID string
    conn   *websocket.Conn
    send   chan any
    hub    *Hub
}

func (c *Client) readPump() {
    log.Printf("[readPump] Starting read pump for user: %s", c.userID)

    defer func() {
        log.Printf("[readPump] Closing connection for user: %s", c.userID)
        c.hub.unregister <- c
        c.conn.Close()
    }()
    c.conn.SetReadLimit(1024)
    c.conn.SetReadDeadline(time.Now().Add(60 * time.Second))
    c.conn.SetPongHandler(func(string) error {
        log.Printf("[readPump] Received pong from user: %s", c.userID)
        c.conn.SetReadDeadline(time.Now().Add(60 * time.Second))
        return nil
    })

    for {
        var msg IncomingMessage
        if err := c.conn.ReadJSON(&msg); err != nil {
            log.Printf("[readPump] Error reading message from user %s: %v", c.userID, err)
            break
        }
        log.Printf("[readPump] Received message from user %s: %+v", c.userID, msg)
        DispatchMessage(c, msg)
    }
}

func (c *Client) writePump() {
    log.Printf("[writePump] Starting write pump for user: %s", c.userID)

    ticker := time.NewTicker(50 * time.Second)
    defer func() {
        log.Printf("[writePump] Closing writer for user: %s", c.userID)
        ticker.Stop()
        c.conn.Close()
    }()
    for {
        select {
        case msg, ok := <-c.send:
            c.conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
            if !ok {
                log.Printf("[writePump] Send channel closed for user: %s", c.userID)
                c.conn.WriteMessage(websocket.CloseMessage, []byte{})
                return
            }
            log.Printf("[writePump] Sending message to user %s: %+v", c.userID, msg)
            if err := c.conn.WriteJSON(msg); err != nil {
                log.Printf("[writePump] Error sending message to user %s: %v", c.userID, err)
                return
            }
        case <-ticker.C:
            c.conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
            if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
                log.Printf("[writePump] Error sending ping to user %s: %v", c.userID, err)
                return
            }
            log.Printf("[writePump] Sent ping to user: %s", c.userID)
        }
    }
}