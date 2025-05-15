package bootstrap

import (
	"messenger/internal/app/service"
	"messenger/internal/infra/adapter/in/rest"
	"messenger/internal/infra/adapter/in/ws"
	"messenger/internal/infra/adapter/out/keycloak"
	"messenger/internal/infra/adapter/out/persistence/repository"
	"messenger/internal/infra/config"

	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
	"gorm.io/gorm"
)

// ServerParams defines the parameters required by the server, including handlers, middleware, and the WebSocket hub.
type ServerParams struct {
	dig.In

	WsHub *ws.Hub
	CorsMiddleware gin.HandlerFunc `name:"cors"`
	JwtMiddleware  gin.HandlerFunc `name:"jwt"`

	MessageHandler *rest.MessageHandler
	KeycloakHandler *rest.KeycloakHandler
	ConversationHandler *rest.ConversationHandler
	ConversationMemberHandler *rest.ConversationMemberHandler
}

// BuildContainer initializes the dependency injection container and registers all dependencies.
func BuildContainer() *dig.Container {
	c := dig.New()
	cfg := config.LoadConfig()
	hub := ws.NewHub()

    // Registering the basic dependencies (database, config, etc.)
	c.Provide(func() *config.Config {
		return cfg
	}) 

	c.Provide(func() *keycloak.KeycloakClient {
		return keycloak.NewKeycloakClient(cfg.Keycloak)
	})

	c.Provide(func() *gorm.DB {
		return config.InitDB(cfg.Database)
	})

    c.Provide(func(cfg *config.Config) gin.HandlerFunc {
		return rest.JWTMiddleware(cfg)
	}, dig.Name("jwt"))

	c.Provide(func(cfg *config.Config) gin.HandlerFunc {
		return config.CORSMiddleware()
	}, dig.Name("cors"))

	// Registering the websocket hub
	c.Provide(func() *ws.Hub {
		return hub
	})

	// Register repositories
	c.Provide(repository.NewConversationMemberRepository)
	c.Provide(repository.NewConversationRepository)
	c.Provide(repository.NewMessageRepository)

	// Register services
	c.Provide(service.NewConversationService)
	c.Provide(service.NewConversationMemberService)
	c.Provide(service.NewMessageService)

    // Register message repo, service, and handler
    c.Provide(repository.NewMessageRepository)
    c.Provide(service.NewMessageService)
    c.Provide(rest.NewMessageHandler)

	c.Provide(rest.NewKeycloakHandler)
	c.Provide(rest.NewConversationHandler)
	c.Provide(rest.NewMessageHandler)
	c.Provide(rest.NewConversationMemberHandler)

	return c
}