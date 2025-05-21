package out

import (
	"messenger/internal/domain"
	"github.com/google/uuid"
)

// ConversationMemberRepository defines the methods for managing conversation members.
type ConversationMemberRepository interface {
	AddMember(member *domain.ConversationMember) error
	RemoveMember(conversationID, userID uuid.UUID) error
	IsMember(conversationID, userID uuid.UUID) (bool, error)
	GetMembers(conversationID uuid.UUID) ([]domain.ConversationMember, error)
}