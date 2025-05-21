package security

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// ExtractUserID retrieves the user ID from the context claims.
func ExtractUserID(c *gin.Context) (uuid.UUID, error) {
	claims, ok := c.Get("claims")
	if !ok {
		return uuid.Nil, errors.New("no claims in context")
	}
	mapClaims := claims.(map[string]interface{})
	sub, ok := mapClaims["sub"].(string)
	if !ok {
		return uuid.Nil, errors.New("invalid subject claim")
	}
	return uuid.Parse(sub)
}

// RequireRole is a middleware that checks if the user has one of the allowed roles.
func RequireRole(allowedRoles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		roleAny, exists := c.Get("role")
		if !exists {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "role missing"})
			return
		}
		role := roleAny.(string)
		for _, allowed := range allowedRoles {
			if role == allowed {
				c.Next()
				return
			}
		}
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "forbidden"})
	}
}
