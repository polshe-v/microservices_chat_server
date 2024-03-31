package service

import (
	"context"

	"github.com/polshe-v/microservices_chat_server/internal/model"
)

// ChatService is the interface for service communication.
type ChatService interface {
	Create(ctx context.Context, chat *model.Chat) (string, error)
	Delete(ctx context.Context, id string) error
}
