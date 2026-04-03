package chat

import (
	"strings"
	"time"

	"github.com/devi-vahid/devichatapp/internal/domain"
)

type MessageRepository interface {
	SaveMessage(msg domain.Message) (domain.Message, error)
}

type Service struct {
	messageRepo MessageRepository
}

type SendMessageInput struct {
	FromUserID int
	ToUserID   int
	Content    string
}

func NewService(messageRepo MessageRepository) *Service {
	return &Service{
		messageRepo: messageRepo,
	}
}

func (s *Service) SendMessage(input SendMessageInput) (domain.Message, error) {
	if input.FromUserID <= 0 {
		return domain.Message{}, domain.ErrInvalidFromUserID
	}

	if input.ToUserID <= 0 {
		return domain.Message{}, domain.ErrInvalidToUserID
	}

	content := strings.TrimSpace(input.Content)
	if content == "" {
		return domain.Message{}, domain.ErrEmptyContent
	}

	msg := domain.Message{
		FromUserID: input.FromUserID,
		ToUserID:   input.ToUserID,
		Content:    content,
		SentAt:     time.Now().UTC(),
	}

	savedMsg, err := s.messageRepo.SaveMessage(msg)
	if err != nil {
		return domain.Message{}, err
	}

	return savedMsg, nil
}
