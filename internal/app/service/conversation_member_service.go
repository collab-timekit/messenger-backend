package service

import (
	"messenger/internal/app/port/out"
	"messenger/internal/domain"
	"github.com/google/uuid"
)

// ConversationMemberService provides methods for managing conversation members.
type ConversationMemberService struct {
	memberRepo out.ConversationMemberRepository
}

// NewConversationMemberService creates a new instance of ConversationMemberService.
func NewConversationMemberService(memberRepo out.ConversationMemberRepository) *ConversationMemberService {
	return &ConversationMemberService{memberRepo: memberRepo}
}

// AddMember adds a user to a conversation.
func (s *ConversationMemberService) AddMember(conversationID, userID uuid.UUID, role string) error {
	member := &domain.ConversationMember{
		ConversationID: conversationID,
		UserID:         userID,
		Role:           role,
	}

	return s.memberRepo.AddMember(member)
}

// RemoveMember removes a user from a conversation.
func (s *ConversationMemberService) RemoveMember(conversationID, userID uuid.UUID) error {
	return s.memberRepo.RemoveMember(conversationID, userID)
}

// GetMembers returns all members of a conversation.
func (s *ConversationMemberService) GetMembers(conversationID uuid.UUID) ([]domain.ConversationMember, error) {
	return s.memberRepo.GetMembers(conversationID)
}

// IsMember checks if a user is a member of a conversation.
func (s *ConversationMemberService) IsMember(conversationID, userID uuid.UUID) (bool, error) {
	return s.memberRepo.IsMember(conversationID, userID)
}