// cmd/chatapi/main.go
package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/devi-vahid/devichatapp/internal/app/chat"
	"github.com/devi-vahid/devichatapp/internal/httpapi"
	"github.com/devi-vahid/devichatapp/internal/store/memory"
)

func main() {
	// create repository and service
	messageRepo := memory.NewMessageRepository()
	chatService := chat.NewService(messageRepo)

	// build HTTP handler and router
	handler := httpapi.NewHandler(chatService)
	mux := http.NewServeMux()
	handler.RegisterRoutes(mux)

	fmt.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
