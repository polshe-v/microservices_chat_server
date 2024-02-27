package chat

import (
	"context"
	"log"

	"github.com/golang/protobuf/ptypes/empty"

	desc "github.com/polshe-v/microservices_chat_server/pkg/chat_v1"
)

// SendMessage is used for sending messages to connected chat.
func (i *Implementation) SendMessage(_ context.Context, req *desc.SendMessageRequest) (*empty.Empty, error) {
	message := req.GetMessage()
	log.Printf("\n%s\nFrom: %s\nText: %s\nTimestamp: %v\n%s", delim, message.GetFrom(), message.GetText(), message.GetTimestamp(), delim)

	return &empty.Empty{}, nil
}
