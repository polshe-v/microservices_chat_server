package chat

import (
	"context"

	"github.com/polshe-v/microservices_chat_server/internal/model"
)

func (s *serv) Create(ctx context.Context, chat *model.Chat) (int64, error) {
	id, err := s.chatRepository.Create(ctx, chat)
	if err != nil {
		return 0, err
	}

	return id, nil
}
