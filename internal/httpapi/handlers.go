package httpapi

import (
	"github.com/devi-vahid/devichatapp/internal/app/chat"
)

type Handler struct {
	ChatService *chat.Service
}

func newHandler(chatService *chat.Service) *Handler {
	return &Handler{
		ChatService: chatService,
	}
}
