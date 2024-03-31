package repository

import (
	"context"

	"github.com/polshe-v/microservices_chat_server/internal/model"
)

// ChatRepository is the interface for user info repository communication.
type ChatRepository interface {
	Create(ctx context.Context, chat *model.Chat) (string, error)
	Delete(ctx context.Context, id string) error
}

// LogRepository is the interface for transaction log repository communication.
type LogRepository interface {
	Log(ctx context.Context, log *model.Log) error
}
