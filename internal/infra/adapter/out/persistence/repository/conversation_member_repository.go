package repository

import (
	"messenger/internal/app/port/out"
	"messenger/internal/domain"
	"messenger/internal/infra/adapter/out/persistence/model"
	"time"

	"github.com/google/uuid"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

// ConversationMemberRepository implements the ConversationMemberRepository interface.
type ConversationMemberRepository struct {
	db *gorm.DB
}

// NewConversationMemberRepository creates a new instance of ConversationMemberRepository.
func NewConversationMemberRepository(db *gorm.DB) out.ConversationMemberRepository {
	return &ConversationMemberRepository{db: db}
}

// AddMember adds a user to a conversation.
func (r *ConversationMemberRepository) AddMember(member *domain.ConversationMember) error {
	var dbMember model.ConversationMemberModel
	_ = copier.Copy(&dbMember, member)

	if dbMember.JoinedAt.IsZero() {
		dbMember.JoinedAt = time.Now()
	}

	return r.db.Create(&dbMember).Error
}

// RemoveMember removes a user from a conversation.
func (r *ConversationMemberRepository) RemoveMember(conversationID, userID uuid.UUID) error {
	return r.db.Delete(&model.ConversationMemberModel{},
		"conversation_id = ? AND user_id = ?", conversationID, userID).Error
}

// IsMember checks if a user is a member of a conversation.
func (r *ConversationMemberRepository) IsMember(conversationID, userID uuid.UUID) (bool, error) {
	var count int64
	err := r.db.Model(&model.ConversationMemberModel{}).
		Where("conversation_id = ? AND user_id = ?", conversationID, userID).
		Count(&count).Error

	return count > 0, err
}

// GetMembers returns all members of a conversation.
func (r *ConversationMemberRepository) GetMembers(conversationID uuid.UUID) ([]domain.ConversationMember, error) {
	var dbMembers []model.ConversationMemberModel

	err := r.db.
		Where("conversation_id = ?", conversationID).
		Find(&dbMembers).Error

	if err != nil {
		return nil, err
	}

	var members []domain.ConversationMember
	_ = copier.Copy(&members, &dbMembers)

	return members, nil
}
