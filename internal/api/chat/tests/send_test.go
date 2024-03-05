package tests

import (
	"context"
	"testing"

	"github.com/gojuno/minimock/v3"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/timestamppb"

	chatAPI "github.com/polshe-v/microservices_chat_server/internal/api/chat"
	"github.com/polshe-v/microservices_chat_server/internal/service"
	serviceMocks "github.com/polshe-v/microservices_chat_server/internal/service/mocks"
	desc "github.com/polshe-v/microservices_chat_server/pkg/chat_v1"
)

func TestSend(t *testing.T) {
	t.Parallel()

	type chatServiceMockFunc func(mc *minimock.Controller) service.ChatService

	type args struct {
		ctx context.Context
		req *desc.SendMessageRequest
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		from      = "from"
		text      = "text"
		timestamp = timestamppb.Now()

		req = &desc.SendMessageRequest{
			Message: &desc.Message{
				From:      from,
				Text:      text,
				Timestamp: timestamp,
			},
		}

		res = &empty.Empty{}
	)

	tests := []struct {
		name            string
		args            args
		want            *empty.Empty
		err             error
		chatServiceMock chatServiceMockFunc
	}{
		{
			name: "success case",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: res,
			err:  nil,
			chatServiceMock: func(mc *minimock.Controller) service.ChatService {
				mock := serviceMocks.NewChatServiceMock(mc)
				return mock
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			chatServiceMock := tt.chatServiceMock(mc)
			api := chatAPI.NewImplementation(chatServiceMock)

			res, err := api.SendMessage(tt.args.ctx, tt.args.req)
			require.Equal(t, tt.err, err)
			require.Equal(t, tt.want, res)
		})
	}
}
