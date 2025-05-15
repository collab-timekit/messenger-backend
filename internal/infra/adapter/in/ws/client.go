package ws

import (
    "github.com/gorilla/websocket"
    "log"
    "time"
)

const (
    pongWait       = 60 * time.Second
    writeWait      = 10 * time.Second
    pingInterval   = 50 * time.Second
    maxMessageSize = 1024
)

// Client represents a WebSocket client connected to the hub.
type Client struct {
    userID string
    conn   *websocket.Conn
    send   chan any
    hub    *Hub
}

func (c *Client) readPump() {
    defer func() {
        c.hub.unregister <- c
        c.conn.Close()
        log.Printf("[readPump] Connection closed for user: %s", c.userID)
    }()

    c.conn.SetReadLimit(maxMessageSize)
    c.conn.SetReadDeadline(time.Now().Add(pongWait))
    c.conn.SetPongHandler(func(string) error {
        c.conn.SetReadDeadline(time.Now().Add(pongWait))
        return nil
    })

    for {
        var msg IncomingMessage
        if err := c.conn.ReadJSON(&msg); err != nil {
            log.Printf("[readPump] Read error from user %s: %v", c.userID, err)
            break
        }
        log.Printf("[readPump] Message from %s: %+v", c.userID, msg)
        DispatchMessage(c, msg)
    }
}

func (c *Client) writePump() {
    ticker := time.NewTicker(pingInterval)
    defer func() {
        ticker.Stop()
        c.conn.Close()
        log.Printf("[writePump] Writer closed for user: %s", c.userID)
    }()

    for {
        select {
        case msg, ok := <-c.send:
            c.conn.SetWriteDeadline(time.Now().Add(writeWait))
            if !ok {
                _ = c.conn.WriteMessage(websocket.CloseMessage, []byte{})
                return
            }
            if err := c.conn.WriteJSON(msg); err != nil {
                log.Printf("[writePump] Write error to user %s: %v", c.userID, err)
                return
            }

        case <-ticker.C:
            c.conn.SetWriteDeadline(time.Now().Add(writeWait))
            if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
                log.Printf("[writePump] Ping error to user %s: %v", c.userID, err)
                return
            }
        }
    }
}