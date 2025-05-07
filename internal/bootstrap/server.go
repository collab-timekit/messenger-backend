package bootstrap

import (
	"messenger/internal/infra/adapter/in/rest"
	"messenger/internal/infra/adapter/in/ws"

	"github.com/gin-gonic/gin"
)

// StartServer initializes the server and starts listening for incoming requests.
func StartServer() error {
    container := buildContainer()

    return container.Invoke(func(
        chatHandler *rest.ChatHandler,
        messageHandler *rest.MessageHandler,
        jwtMiddleware gin.HandlerFunc,
        wsHub *ws.Hub,
    ) error {
        r := gin.Default()

        api := r.Group("/api")
        // api.Use(jwtMiddleware)

        chatHandler.RegisterRoutes(api)
        messageHandler.RegisterRoutes(api)

        r.GET("/ws", ws.ServeWs(wsHub))

        return r.Run(":8083")
    })
}