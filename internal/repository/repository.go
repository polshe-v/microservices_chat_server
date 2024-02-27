package repository

import (
	"context"

	"github.com/polshe-v/microservices_chat_server/internal/model"
)

// ChatRepository is the interface for repository communication.
type ChatRepository interface {
	Create(ctx context.Context, chat *model.Chat) (int64, error)
	Delete(ctx context.Context, id int64) error
}