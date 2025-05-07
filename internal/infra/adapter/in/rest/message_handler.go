package rest

import (
	"net/http"
	"strconv"

	"messenger/internal/app/service"
	"github.com/gin-gonic/gin"
)

type MessageHandler struct {
	service *service.MessageService
}

func NewMessageHandler(service *service.MessageService) *MessageHandler {
	return &MessageHandler{service: service}
}

func (h *MessageHandler) RegisterRoutes(rg *gin.RouterGroup) {
	rg.GET("/chats/:id/messages", h.ListMessages)
	rg.POST("/chats/:id/messages", h.CreateMessage)
}

func (h *MessageHandler) ListMessages(c *gin.Context) {
	chatID := c.Param("id")
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "20"))
	offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))

	messages, err := h.service.GetMessages(chatID, limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, messages)
}

func (h *MessageHandler) CreateMessage(c *gin.Context) {
	chatID := c.Param("id")

	var req struct {
		SenderID string `json:"senderId"`
		Content  string `json:"content"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	message, err := h.service.SendMessage(chatID, req.SenderID, req.Content)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, message)
}