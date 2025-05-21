package ws

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
    CheckOrigin: func(r *http.Request) bool { return true },
}

// ServeWs upgrades the HTTP connection to a WebSocket connection and registers the client to the hub.
func ServeWs(c *gin.Context, hub *Hub) {
    log.Println("Received WebSocket upgrade request")
    userID := c.DefaultQuery("userId", "unknown")
    log.Printf("Incoming WebSocket connection for userId: %s\n", userID)

    conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
    if err != nil {
        log.Printf("Error during connection upgrade: %v\n", err)
        return
    }

    log.Printf("Connection upgraded to WebSocket for userId: %s\n", userID)

    client := &Client{
        userID: userID,
        conn:   conn,
        send:   make(chan any, 256),
        hub:    hub,
    }

    log.Printf("Registering new client with userId: %s\n", userID)
    hub.register <- client

    log.Printf("Starting communication pumps for userId: %s\n", userID)
    go client.writePump()
    go client.readPump()

    log.Printf("WebSocket connection established for userId: %s\n", userID)
}