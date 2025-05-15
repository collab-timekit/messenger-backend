package bootstrap

import (
	"messenger/internal/infra/adapter/in/ws"
	"github.com/gin-gonic/gin"
)

// StartServer initializes the server and starts listening for incoming requests.
func StartServer() error {
    container := BuildContainer()

    return container.Invoke(func(p ServerParams) error {
        r := gin.Default()

        // Start the WebSocket hub in a separate goroutine
        go p.WsHub.Run()

        api := r.Group("/api")
        // api.Use(jwtMiddleware)
        api.Use(p.CorsMiddleware)

        p.ConversationMemberHandler.RegisterRoutes(api)
        p.ConversationHandler.RegisterRoutes(api)
        p.KeycloakHandler.RegisterRoutes(api)
        p.MessageHandler.RegisterRoutes(api)

        r.GET("/ws", func(c *gin.Context) {
			ws.ServeWs(c, p.WsHub)
		})

        return r.Run(":8083")
    })
}