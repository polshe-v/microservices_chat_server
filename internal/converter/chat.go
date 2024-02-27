package converter

import (
	"github.com/polshe-v/microservices_chat_server/internal/model"
	desc "github.com/polshe-v/microservices_chat_server/pkg/chat_v1"
)

// ToChatFromService converts service layer model to structure of API layer.
func ToChatFromService(chat *model.Chat) *desc.Chat {
	return &desc.Chat{
		Usernames: chat.Usernames,
	}
}

// ToChatFromDesc converts structure of API layer to service layer model.
func ToChatFromDesc(chat *desc.Chat) *model.Chat {
	return &model.Chat{
		Usernames: chat.Usernames,
	}
}
