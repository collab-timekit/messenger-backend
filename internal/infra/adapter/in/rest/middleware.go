package rest

import (
	"context"
	"net/http"
	"strings"

	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/gin-gonic/gin"
	"messenger/internal/infra/config"
)

// JWTMiddleware is a middleware that verifies JWT tokens using Keycloak.
func JWTMiddleware(cfg *config.Config) gin.HandlerFunc {
	provider, err := oidc.NewProvider(context.Background(), cfg.Keycloak.Issuer)
	if err != nil {
		panic(err)
	}

	verifier := provider.Verifier(&oidc.Config{ClientID: cfg.Keycloak.ClientID, SkipIssuerCheck: true}) //TODO: remove SkipIssuerCheck in production

	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if !strings.HasPrefix(authHeader, "Bearer ") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Missing or invalid token"})
			return
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")
		idToken, err := verifier.Verify(c.Request.Context(), token)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token: " + err.Error()})
			return
		}

		var claims map[string]interface{}
		if err := idToken.Claims(&claims); err == nil {
			c.Set("claims", claims)
		}

		c.Next()
	}
}