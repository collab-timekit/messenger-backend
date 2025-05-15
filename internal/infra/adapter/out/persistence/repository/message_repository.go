package repository

import (
	"messenger/internal/app/port/out"
	"messenger/internal/domain"
	"messenger/internal/infra/adapter/out/persistence/model"
	"time"

	"github.com/google/uuid"
	"github.com/jinzhu/copier"
	"golang.org/x/exp/slices"
	"gorm.io/gorm"
)

type MessageRepository struct {
	db *gorm.DB
}

func NewMessageRepository(db *gorm.DB) out.MessageRepository {
	return &MessageRepository{db: db}
}

// SendMessage saves a new message to the database.
func (r *MessageRepository) SendMessage(msg *domain.Message) error {
	var dbMsg model.MessageModel
	_ = copier.Copy(&dbMsg, msg)
	return r.db.Create(&dbMsg).Error
}

// GetMessages returns messages from a conversation with pagination.
// 'before' is optional. If present, fetch messages older than that timestamp.
// Results are returned ascending (oldest to newest).
func (r *MessageRepository) GetMessages(conversationID uuid.UUID, limit int, before *time.Time) ([]domain.Message, error) {
	var dbMessages []model.MessageModel

	query := r.db.
		Where("conversation_id = ?", conversationID).
		Order("created_at DESC").
		Limit(limit)

	if before != nil {
		query = query.Where("created_at < ?", *before)
	}

	err := query.Find(&dbMessages).Error
	if err != nil {
		return nil, err
	}

	var messages []domain.Message
	_ = copier.Copy(&messages, &dbMessages)

	// Because we query DESC, reverse to return ASC (oldest → newest)
	slices.Reverse(messages)

	return messages, nil
}

func (r *ConversationRepository) GetByID(id uuid.UUID) (*domain.Conversation, error) {
	var modelConv model.ConversationModel
	if err := r.db.Preload("Members").First(&modelConv, "id = ?", id).Error; err != nil {
		return nil, err
	}

	var conv domain.Conversation
	_ = copier.Copy(&conv, &modelConv)
	return &conv, nil
}

// EditMessage updates the content of a message (optional feature).
func (r *MessageRepository) EditMessage(messageID uuid.UUID, newContent string) error {
	return r.db.Model(&model.MessageModel{}).
		Where("id = ?", messageID).
		Updates(map[string]interface{}{
			"content": newContent,
			"edited":  true,
		}).Error
}

// DeleteMessage removes a message by ID (optional feature).
func (r *MessageRepository) DeleteMessage(messageID uuid.UUID) error {
	return r.db.Delete(&model.MessageModel{}, "id = ?", messageID).Error
}