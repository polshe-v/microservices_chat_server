package chat

import (
	"github.com/polshe-v/microservices_chat_server/internal/client/db"
	"github.com/polshe-v/microservices_chat_server/internal/repository"
	"github.com/polshe-v/microservices_chat_server/internal/service"
)

type serv struct {
	chatRepository repository.ChatRepository
	logRepository  repository.LogRepository
	txManager      db.TxManager
}

// NewService creates new object of service layer.
func NewService(chatRepository repository.ChatRepository, logRepository repository.LogRepository, txManager db.TxManager) service.ChatService {
	return &serv{
		chatRepository: chatRepository,
		logRepository:  logRepository,
		txManager:      txManager,
	}
}
