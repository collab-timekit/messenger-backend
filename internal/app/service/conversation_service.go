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

// GetOrCreatePrivateConversation checks if a private chat exists, returns its ID or creates one.
func (s *ConversationService) GetOrCreatePrivateConversation(userID, recipientID uuid.UUID) (uuid.UUID, error) {
	// Check if the conversation already exists
	existingID, err := s.conversationRepo.FindPrivateConversationIDBetweenUsers(userID, recipientID)
	if err != nil {
		return uuid.Nil, err
	}
	if existingID != nil {
		return *existingID, nil
	}

	// Create new conversation
	conv := &domain.Conversation{
		ID:   uuid.New(),
		Type: domain.Private,
		Members: []domain.ConversationMember{
			{UserID: userID, Role: "member"},
			{UserID: recipientID, Role: "member"},
		},
	}
	if err := s.conversationRepo.Create(conv); err != nil {
		return uuid.Nil, err
	}
	return conv.ID, nil
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

// CreateConversation creates a new conversation (channel or group).
func (s *ConversationService) CreateConversation(conv *domain.Conversation) error {
	return s.conversationRepo.Create(conv)
}