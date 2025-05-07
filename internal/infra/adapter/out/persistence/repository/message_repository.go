package repository

import (
	"messenger/internal/app/port/out"
	"messenger/internal/domain"
	"messenger/internal/infra/adapter/out/persistence/model"

	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

type MessageRepository struct {
	db *gorm.DB
}

func NewMessageRepository(db *gorm.DB) out.MessageRepository {
	return &MessageRepository{db: db}
}

func (r *MessageRepository) FindByChatID(chatID string, limit, offset int) ([]domain.Message, error) {
	var models []model.MessageModel
	if err := r.db.
		Where("chat_id = ?", chatID).
		Order("created_at DESC").
		Limit(limit).
		Offset(offset).
		Find(&models).Error; err != nil {
		return nil, err
	}

	// Mapowanie z modelu na domenę z użyciem copier
	var result []domain.Message
	if err := copier.Copy(&result, &models); err != nil {
		return nil, err
	}

	return result, nil
}

func (r *MessageRepository) Save(msg *domain.Message) error {
	// Mapowanie z domeny na model przed zapisem do bazy danych
	var model model.MessageModel
	if err := copier.Copy(&model, msg); err != nil {
		return err
	}

	return r.db.Create(&model).Error
}