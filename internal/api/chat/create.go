package chat

import (
	"context"
	"log"
	"strings"

	"github.com/polshe-v/microservices_chat_server/internal/converter"
	desc "github.com/polshe-v/microservices_chat_server/pkg/chat_v1"
)

// Create is used for creating new chat.
func (i *Implementation) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	chat := req.GetChat()
	log.Printf("\n%s\nUsernames: %v\n%s", delim, strings.Join(chat.GetUsernames(), ", "), delim)

	id, err := i.chatService.Create(ctx, converter.ToChatFromDesc(chat))
	if err != nil {
		return nil, err
	}

	log.Printf("Created chat with id: %d", id)

	return &desc.CreateResponse{
		Id: id,
	}, nil
}
