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
    // Log the incoming request
    userID := c.DefaultQuery("userId", "unknown")
    log.Printf("Incoming WebSocket connection for userId: %s\n", userID)

    // Try upgrading the connection to WebSocket
    conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
    if err != nil {
        log.Printf("Error during connection upgrade: %v\n", err)
        return
    }

    // Log successful connection upgrade
    log.Printf("Connection upgraded to WebSocket for userId: %s\n", userID)

    // Create the client and add to the hub
    client := &Client{
        userID: userID,
        conn:   conn,
        send:   make(chan any, 256),
        hub:    hub,
    }

    // Log when the client is registered in the hub
    log.Printf("Registering new client with userId: %s\n", userID)
    hub.register <- client

    // Start the client communication loops (writePump and readPump)
    log.Printf("Starting communication pumps for userId: %s\n", userID)
    go client.writePump()
    go client.readPump()

    log.Printf("WebSocket connection established for userId: %s\n", userID)
}