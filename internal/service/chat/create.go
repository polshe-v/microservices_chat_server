package chat

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/polshe-v/microservices_chat_server/internal/model"
)

func (s *serv) Create(ctx context.Context, chat *model.Chat) (int64, error) {
	var id int64

	err := s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		var errTx error
		id, errTx = s.chatRepository.Create(ctx, chat)
		if errTx != nil {
			return errTx
		}

		errTx = s.logRepository.Log(ctx, &model.Log{
			Log: fmt.Sprintf("Created chat with id: %d", id),
		})
		if errTx != nil {
			return errTx
		}

		return nil
	})

	if err != nil {
		log.Print(err)
		return 0, errors.New("failed to create chat")
	}

	return id, nil
}
