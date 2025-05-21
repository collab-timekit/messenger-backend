package rest

import (
	"messenger/internal/app/port/in"
	"messenger/internal/infra/security"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// MessageHandler handles HTTP requests related to messages.
type MessageHandler struct {
	useCase in.MessageUseCase
}

// NewMessageHandler creates a new instance of MessageHandler.
func NewMessageHandler(useCase in.MessageUseCase) *MessageHandler {
	return &MessageHandler{useCase: useCase}
}

// RegisterRoutes registers the routes for message operations.
func (h *MessageHandler) RegisterRoutes(rg *gin.RouterGroup) {
	rg.POST("/messages", h.CreateMessage)
	rg.GET("/messages/:conversation_id", h.GetMessages)
	rg.PUT("/messages/:message_id", h.EditMessage)
	rg.DELETE("/messages/:message_id", h.DeleteMessage)
}

// CreateMessage creates a new message in a conversation.
func (h *MessageHandler) CreateMessage(c *gin.Context) {
	userID, err := security.ExtractUserID(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	var req struct {
		ConversationID uuid.UUID `json:"conversation_id" binding:"required"`
		Content        string    `json:"content" binding:"required"`
	}

	bindErr := c.ShouldBindJSON(&req)
	if bindErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": bindErr.Error()})
		return
	}

	err = h.useCase.CreateMessage(req.ConversationID, userID, req.Content)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not send message"})
		return
	}

	c.Status(http.StatusCreated)
}

// GetMessages retrieves messages for a given conversation with optional pagination.
func (h *MessageHandler) GetMessages(c *gin.Context) {
	conversationID, err := uuid.Parse(c.Param("conversation_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid conversation id"})
		return
	}

	limit := 50
	if l := c.Query("limit"); l != "" {
		if parsed, parseErr := strconv.Atoi(l); parseErr == nil {
			limit = parsed
		}
	}

	var before *time.Time
	if b := c.Query("before"); b != "" {
		if parsedTime, parseErr := time.Parse(time.RFC3339, b); parseErr == nil {
			before = &parsedTime
		}
	}

	messages, err := h.useCase.GetMessagesForConversation(conversationID, limit, before)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not get messages"})
		return
	}

	c.JSON(http.StatusOK, messages)
}

// EditMessage allows editing an existing message.
func (h *MessageHandler) EditMessage(c *gin.Context) {
	_, err := security.ExtractUserID(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	messageID, err := uuid.Parse(c.Param("message_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid message id"})
		return
	}

	var req struct {
		NewContent string `json:"new_content" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.useCase.EditMessage(messageID, req.NewContent); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not edit message"})
		return
	}

	c.Status(http.StatusOK)
}

// DeleteMessage deletes a message from a conversation.
func (h *MessageHandler) DeleteMessage(c *gin.Context) {
	_, err := security.ExtractUserID(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	messageID, err := uuid.Parse(c.Param("message_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid message id"})
		return
	}

	if err := h.useCase.DeleteMessage(messageID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not delete message"})
		return
	}

	c.Status(http.StatusOK)
}