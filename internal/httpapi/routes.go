package httpapi

import "net/http"

// RegisterRoutes attaches all API endpoints to the mux.
func (h *Handler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/api/v1/messages/send", h.SendMessage)
}
