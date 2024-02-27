package chat

import (
	"github.com/polshe-v/microservices_chat_server/internal/repository"
	"github.com/polshe-v/microservices_chat_server/internal/service"
)

type serv struct {
	chatRepository repository.ChatRepository
}

// NewService creates new object of service layer.
func NewService(chatRepository repository.ChatRepository) service.ChatService {
	return &serv{
		chatRepository: chatRepository,
	}
}
