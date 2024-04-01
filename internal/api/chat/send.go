package chat

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"

	"github.com/polshe-v/microservices_chat_server/internal/converter"
	desc "github.com/polshe-v/microservices_chat_server/pkg/chat_v1"
)

// SendMessage is used for sending messages to connected chat.
func (i *Implementation) SendMessage(ctx context.Context, req *desc.SendMessageRequest) (*empty.Empty, error) {
	err := i.chatService.SendMessage(ctx, req.GetChatId(), converter.ToMessageFromDesc(req.GetMessage()))
	if err != nil {
		return nil, err
	}

	return &empty.Empty{}, nil
}
