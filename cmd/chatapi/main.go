package main

import (
	"fmt"

	"github.com/devi-vahid/devichatapp/internal/app/chat"
	"github.com/devi-vahid/devichatapp/internal/store/memory"
)

func main() {
	messageRepo := memory.NewMessageRepository()
	chatService := chat.NewService(messageRepo)

	fmt.Printf("chat service initialized: %+v\n", chatService)
}