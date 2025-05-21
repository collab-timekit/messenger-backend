package rest

import (
	"messenger/internal/app/port/in"
	"messenger/internal/domain"
	"messenger/internal/infra/security"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// ConversationHandler handles HTTP requests related to conversations.
type ConversationHandler struct {
	useCase in.ConversationUseCase
}

// NewConversationHandler creates a new instance of ConversationHandler.
func NewConversationHandler(useCase in.ConversationUseCase) *ConversationHandler {
	return &ConversationHandler{useCase: useCase}
}

// RegisterRoutes registers the routes for conversation operations.
func (h *ConversationHandler) RegisterRoutes(rg *gin.RouterGroup) {
	rg.POST("/chats", h.CreateChat)
	rg.GET("/chats", h.ListChats)
	rg.GET("/chats/:id", h.GetChatByID)
}

// CreateChat creates a new conversation.
func (h *ConversationHandler) CreateChat(c *gin.Context) {
	userID, err := security.ExtractUserID(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	var req struct {
		Type    domain.ConversationType   `json:"type" binding:"required"`
		Name    *string  `json:"name"`
		Members []string `json:"members"` // UUID w stringach
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	conv := &domain.Conversation{
		ID:   uuid.New(),
		Type: req.Type,
		Name: req.Name,
		Members: []domain.ConversationMember{
			{UserID: userID, Role: "admin"}, // zakładając, że autor jest adminem
		},
	}

	for _, idStr := range req.Members {
		memberID, err := uuid.Parse(idStr)
		if err != nil || memberID == userID {
			continue
		}
		conv.Members = append(conv.Members, domain.ConversationMember{
			UserID: memberID,
			Role:   "member",
		})
	}

	if err := h.useCase.CreateConversation(conv); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not create chat"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": conv.ID})
}

// ListChats lists all conversations for the authenticated user.
func (h *ConversationHandler) ListChats(c *gin.Context) {
	userID, err := security.ExtractUserID(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	convs, err := h.useCase.GetUserConversations(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not fetch chats"})
		return
	}

	c.JSON(http.StatusOK, convs)
}

// GetChatByID retrieves a conversation by its ID.
func (h *ConversationHandler) GetChatByID(c *gin.Context) {
	userID, err := security.ExtractUserID(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	convIDStr := c.Param("id")
	convID, err := uuid.Parse(convIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid conversation ID"})
		return
	}

	conv, err := h.useCase.GetConversationByID(convID, userID)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "you are not a member of this conversation"})
		return
	}

	c.JSON(http.StatusOK, conv)
}
