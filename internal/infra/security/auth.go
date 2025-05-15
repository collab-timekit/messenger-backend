package security

import (
	"errors"

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
