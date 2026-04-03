package memory

import (
	"sync"

	"github.com/devi-vahid/devichatapp/internal/domain"
)

type MessageRepository struct {
	mu       sync.Mutex
	messages []domain.Message
	nextID   int
}

func NewMessageRepository() *MessageRepository {
	return &MessageRepository{
		messages: make([]domain.Message, 0),
		nextID:   1,
	}
}

func (r *MessageRepository) SaveMessage(msg domain.Message) (domain.Message, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	msg.ID = r.nextID
	r.nextID++

	r.messages = append(r.messages, msg)

	return msg, nil
}
