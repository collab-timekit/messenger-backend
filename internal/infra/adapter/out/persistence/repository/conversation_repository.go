package repository

import (
	"messenger/internal/app/port/out"
	"messenger/internal/domain"
	"messenger/internal/infra/adapter/out/persistence/model"

	"github.com/google/uuid"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

// ConversationRepository implements the ConversationRepository interface.
type ConversationRepository struct {
	db *gorm.DB
}

// NewConversationRepository creates a new instance of ConversationRepository.
func NewConversationRepository(db *gorm.DB) out.ConversationRepository {
	return &ConversationRepository{db: db}
}

// GetAllByUserID retrieves all conversations for a given user ID.
func (r *ConversationRepository) GetAllByUserID(userID uuid.UUID) ([]domain.Conversation, error) {
	var dbConvs []model.ConversationModel
	err := r.db.
		Joins("JOIN conversation_members ON conversation_members.conversation_id = conversations.id").
		Where("conversation_members.user_id = ?", userID).
		Preload("Members").
		Preload("Messages").
		Find(&dbConvs).Error
	if err != nil {
		return nil, err
	}

	var result []domain.Conversation
	for _, dbConv := range dbConvs {
		var conv domain.Conversation
		_ = copier.Copy(&conv, &dbConv)

		conv.Members = make([]domain.ConversationMember, len(dbConv.Members))
		_ = copier.Copy(&conv.Members, &dbConv.Members)

		conv.Messages = make([]domain.Message, len(dbConv.Messages))
		_ = copier.Copy(&conv.Messages, &dbConv.Messages)

		result = append(result, conv)
	}

	return result, nil
}

// Create creates a new conversation in the database.
func (r *ConversationRepository) Create(conv *domain.Conversation) error {
	var dbConv model.ConversationModel
	_ = copier.Copy(&dbConv, conv)

	_ = copier.Copy(&dbConv.Members, &conv.Members)
	_ = copier.Copy(&dbConv.Messages, &conv.Messages)

	return r.db.Create(&dbConv).Error
}

// FindPrivateConversationIDBetweenUsers returns the ID of a private conversation between two users if it exists.
func (r *ConversationRepository) FindPrivateConversationIDBetweenUsers(user1, user2 uuid.UUID) (*uuid.UUID, error) {
	var convID uuid.UUID
	err := r.db.Raw(`
		SELECT c.id
		FROM conversations c
		JOIN conversation_members cm ON cm.conversation_id = c.id
		WHERE c.type = 'PRIVATE'
		AND cm.user_id IN (?, ?)
		GROUP BY c.id
		HAVING COUNT(DISTINCT cm.user_id) = 2
		LIMIT 1
	`, user1, user2).Scan(&convID).Error

	if err != nil {
		return nil, err
	}
	if convID == uuid.Nil {
		return nil, nil
	}
	return &convID, nil
}