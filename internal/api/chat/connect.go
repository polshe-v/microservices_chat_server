package chat

import (
	"github.com/polshe-v/microservices_chat_server/internal/converter"
	desc "github.com/polshe-v/microservices_chat_server/pkg/chat_v1"
)

// Connect is used for connecting to a chat.
func (i *Implementation) Connect(req *desc.ConnectRequest, stream desc.ChatV1_ConnectServer) error {
	return i.chatService.Connect(req.GetChatId(), req.GetUsername(), converter.ToStreamFromDesc(stream))
}
