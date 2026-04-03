package chat

import "github.com/devi-vahid/devichatapp/internal/domain"

type MessageRepository interface {
	SaveMessage(msg domain.Message) (domain.Message, error)
}