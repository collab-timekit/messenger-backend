package bootstrap

import (
	"messenger/internal/app/service"
	"messenger/internal/infra/adapter/in/rest"
	"messenger/internal/infra/adapter/in/ws"
	"messenger/internal/infra/adapter/out/persistence/repository"
	"messenger/internal/infra/config"

	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
	"gorm.io/gorm"
)

func buildContainer() *dig.Container {
	c := dig.New()
	cfg := config.LoadConfig()
	hub := ws.NewHub()

    // Registering the basic dependencies (database, config, etc.)
	c.Provide(func() *config.Config {
		return cfg
	})

	c.Provide(func() *gorm.DB {
		return config.InitDB(cfg.Database)
	})

    c.Provide(func(cfg *config.Config) gin.HandlerFunc {
		return rest.JWTMiddleware(cfg)
	})

	// Registering the websocket hub
	c.Provide(func() *ws.Hub {
		return hub
	})

    /// Registering repositories and services

    // Register chat repo, service, and handler
	c.Provide(repository.NewChatRepository)
	c.Provide(service.NewChatService)
	c.Provide(rest.NewChatHandler)

    // Register message repo, service, and handler
    c.Provide(repository.NewMessageRepository)
    c.Provide(service.NewMessageService)
    c.Provide(rest.NewMessageHandler)

	return c
}