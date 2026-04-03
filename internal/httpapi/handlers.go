package httpapi

import (
	"encoding/json"
	"net/http"

	"github.com/devi-vahid/devichatapp/internal/app/chat"
	"github.com/devi-vahid/devichatapp/internal/domain"
)

type Handler struct {
	chatService *chat.Service
}

func NewHandler(chatService *chat.Service) *Handler {
	return &Handler{
		chatService: chatService,
	}
}

func (h *Handler) SendMessage(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var input chat.SendMessageInput
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&input); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	msg, err := h.chatService.SendMessage(input)
	if err != nil {
		switch err {
		case domain.ErrInvalidFromUserID, domain.ErrInvalidToUserID, domain.ErrEmptyContent:
			http.Error(w, err.Error(), http.StatusBadRequest)
		default:
			http.Error(w, "internal server error", http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(msg)
}
