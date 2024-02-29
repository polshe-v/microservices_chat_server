package chat

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"

	desc "github.com/polshe-v/microservices_chat_server/pkg/chat_v1"
)

// SendMessage is used for sending messages to connected chat.
func (i *Implementation) SendMessage(_ context.Context, _ *desc.SendMessageRequest) (*empty.Empty, error) {
	return &empty.Empty{}, nil
}
