package repository

import (
	"messenger/internal/app/port/out"
	"messenger/internal/domain"
	"messenger/internal/infra/adapter/out/persistence/model"

	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

type ChatRepository struct {
	db *gorm.DB
}

func NewChatRepository(db *gorm.DB) out.ChatRepository {
	return &ChatRepository{db: db}
}

// Save zapisuje Chat w bazie danych
func (r *ChatRepository) Save(c *domain.Chat) error {
	chatModel := model.ChatModel{}
	err := copier.Copy(&chatModel, c) // Użycie copier do mapowania
	if err != nil {
		return err
	}
	return r.db.Create(&chatModel).Error
}

// FindAll znajduje wszystkie czaty w bazie danych
func (r *ChatRepository) FindAll() ([]*domain.Chat, error) {
	var models []model.ChatModel
	if err := r.db.Find(&models).Error; err != nil {
		return nil, err
	}

	result := make([]*domain.Chat, 0, len(models))
	for _, m := range models {
		chat := &domain.Chat{}
		err := copier.Copy(chat, &m) // Użycie copier do mapowania
		if err != nil {
			return nil, err
		}
		result = append(result, chat)
	}
	return result, nil
}

// FindByID znajduje czat po ID
func (r *ChatRepository) FindByID(id string) (*domain.Chat, error) {
	var m model.ChatModel
	if err := r.db.First(&m, "id = ?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	chat := &domain.Chat{}
	err := copier.Copy(chat, &m) // Użycie copier do mapowania
	if err != nil {
		return nil, err
	}
	return chat, nil
}

// GetMessages zwraca wiadomości związane z czatem
func (r *ChatRepository) GetMessages(chatID string, limit, offset int) ([]model.MessageModel, error) {
	var messages []model.MessageModel
	err := r.db.Where("chat_id = ?", chatID).
		Order("created_at ASC").
		Limit(limit).
		Offset(offset).
		Find(&messages).Error
	return messages, err
}