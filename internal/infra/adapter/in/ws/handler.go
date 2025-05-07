package ws

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
    CheckOrigin: func(r *http.Request) bool { return true },
}

// ServeWs upgrades the HTTP connection to a WebSocket connection and registers the client to the hub.
func ServeWs(hub *Hub) gin.HandlerFunc {
    return func(c *gin.Context) {
        userID := c.Query("userId")

        conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
        if err != nil {
            return
        }

        client := &Client{
			userID: userID,
			conn:   conn,
            send:   make(chan any, 256),
			hub:    hub,
		}		

        hub.register <- client

        go client.writePump()
        go client.readPump()
    }
}