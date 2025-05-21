package bootstrap

import (
	"messenger/internal/infra/adapter/in/ws"
	"github.com/gin-gonic/gin"
)

func StartServer() error {
    container := BuildContainer()

    return container.Invoke(func(p ServerParams) error {
        r := gin.Default()

        r.Use(p.CorsMiddleware)

        api := r.Group("/api")
        api.Use(p.JwtMiddleware)

        go p.WsHub.Run()
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