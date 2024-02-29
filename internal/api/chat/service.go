package chat

import (
	"github.com/polshe-v/microservices_chat_server/internal/service"
	desc "github.com/polshe-v/microservices_chat_server/pkg/chat_v1"
)

// Implementation structure describes API layer.
type Implementation struct {
	desc.UnimplementedChatV1Server
	chatService service.ChatService
}

// NewImplementation creates new object of API layer.
func NewImplementation(chatService service.ChatService) *Implementation {
	return &Implementation{
		chatService: chatService,
	}
}
