package service

import (
	"errors"
	"messenger/internal/app/port/in"
	"messenger/internal/app/port/out"
	"messenger/internal/domain"

	"github.com/google/uuid"
)

// ConversationService provides methods for working with conversations.
type ConversationService struct {
	conversationRepo out.ConversationRepository
	memberRepo       out.ConversationMemberRepository
}

// NewConversationService creates a new instance of ConversationService.
func NewConversationService(
	conversationRepo out.ConversationRepository,
	memberRepo out.ConversationMemberRepository,
) in.ConversationUseCase {
	return &ConversationService{
		conversationRepo: conversationRepo,
		memberRepo:       memberRepo,
	}
}

// GetUserConversations returns all conversations for a given user.
func (s *ConversationService) GetUserConversations(userID uuid.UUID) ([]domain.Conversation, error) {
	return s.conversationRepo.GetAllByUserID(userID)
}

// CreateConversation creates a new conversation.
func (s *ConversationService) CreateConversation(conv *domain.Conversation) error {
	return s.conversationRepo.Create(conv)
}


// GetConversationByID returns a conversation by its ID if the user is a member.
func (s *ConversationService) GetConversationByID(conversationID, userID uuid.UUID) (*domain.Conversation, error) {
	isMember, err := s.memberRepo.IsMember(conversationID, userID)
	if err != nil {
		return nil, err
	}
	if !isMember {
		return nil, errors.New("unauthorized: not a member")
	}

	return s.conversationRepo.GetByID(conversationID)
}