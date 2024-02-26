package repository

import (
	"context"

	desc "github.com/polshe-v/microservices_chat_server/pkg/chat_v1"
)

// ChatRepository is the interface for repository communication.
type ChatRepository interface {
	Create(ctx context.Context, chat *desc.Chat) (int64, error)
	Delete(ctx context.Context, id int64) error
}
