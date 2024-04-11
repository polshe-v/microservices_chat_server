package chat

import (
	"sync"

	"github.com/polshe-v/microservices_chat_server/internal/model"
	"github.com/polshe-v/microservices_chat_server/internal/repository"
	"github.com/polshe-v/microservices_chat_server/internal/service"
	"github.com/polshe-v/microservices_common/pkg/db"
)

const messagesBuffer = 100

type serv struct {
	chatRepository     repository.ChatRepository
	messagesRepository repository.MessagesRepository
	logRepository      repository.LogRepository
	txManager          db.TxManager

	channels   map[string]chan *model.Message
	mxChannels sync.RWMutex

	chats  map[string]*chat
	mxChat sync.RWMutex
}

type chat struct {
	streams map[string]model.Stream
	m       sync.RWMutex
}

// NewService creates new object of service layer.
func NewService(chatRepository repository.ChatRepository, messagesRepository repository.MessagesRepository, logRepository repository.LogRepository, txManager db.TxManager) service.ChatService {
	return &serv{
		chatRepository:     chatRepository,
		messagesRepository: messagesRepository,
		logRepository:      logRepository,
		txManager:          txManager,
		chats:              make(map[string]*chat),
		channels:           make(map[string]chan *model.Message),
	}
}
