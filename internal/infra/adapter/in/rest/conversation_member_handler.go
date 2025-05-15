package rest

import (
	"messenger/internal/app/port/in"
	"messenger/internal/infra/security"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ConversationMemberHandler struct {
	useCase in.ConversationMemberUseCase
}

func NewConversationMemberHandler(useCase in.ConversationMemberUseCase) *ConversationMemberHandler {
	return &ConversationMemberHandler{useCase: useCase}
}

func (h *ConversationMemberHandler) RegisterRoutes(rg *gin.RouterGroup) {
	rg.GET("/chats/:id/members", h.ListMembers)
	rg.POST("/chats/:id/members", h.AddMember)
	rg.DELETE("/chats/:id/members/:userID", h.RemoveMember)
}

func (h *ConversationMemberHandler) ListMembers(c *gin.Context) {
	userID, err := security.ExtractUserID(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	convID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid conversation id"})
		return
	}

	isMember, err := h.useCase.IsMember(convID, userID)
	if err != nil || !isMember {
		c.JSON(http.StatusForbidden, gin.H{"error": "not a member"})
		return
	}

	members, err := h.useCase.GetMembers(convID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get members"})
		return
	}

	c.JSON(http.StatusOK, members)
}

func (h *ConversationMemberHandler) AddMember(c *gin.Context) {
	var req struct {
		UserID string `json:"user_id" binding:"required"`
		Role   string `json:"role" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid body"})
		return
	}

	convID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid conversation id"})
		return
	}

	userID, err := uuid.Parse(req.UserID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
		return
	}

	if err := h.useCase.AddMember(convID, userID, req.Role); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to add member"})
		return
	}

	c.Status(http.StatusCreated)
}

func (h *ConversationMemberHandler) RemoveMember(c *gin.Context) {
	convID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid conversation id"})
		return
	}

	userID, err := uuid.Parse(c.Param("userID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
		return
	}

	// Możesz dodać sprawdzenie, czy usuwający ma uprawnienia
	if err := h.useCase.RemoveMember(convID, userID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to remove member"})
		return
	}

	c.Status(http.StatusNoContent)
}
