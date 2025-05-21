package rest

import (
	"net/http"

	"messenger/internal/infra/adapter/out/keycloak"

	"github.com/gin-gonic/gin"
)

// KeycloakHandler handles requests related to Keycloak operations.
type KeycloakHandler struct {
	Client *keycloak.KeycloakClient
}

// NewKeycloakHandler creates a new instance of KeycloakHandler.
func NewKeycloakHandler(client *keycloak.KeycloakClient) *KeycloakHandler {
	return &KeycloakHandler{Client: client}
}

// RegisterRoutes registers the routes for Keycloak operations.
func (h *KeycloakHandler) RegisterRoutes(rg *gin.RouterGroup) {
	rg.GET("/keycloak/users", h.GetUsers)
}

// GetUsers retrieves users from Keycloak based on the provided filter.
func (h *KeycloakHandler) GetUsers(c *gin.Context) {
	filter := c.Query("filter")
	if filter == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "filter query param required"})
		return
	}

	users, err := h.Client.GetUsers(filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}
