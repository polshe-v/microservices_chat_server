package tests

import (
	"context"
	"fmt"
	"testing"

	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"

	chatAPI "github.com/polshe-v/microservices_chat_server/internal/api/chat"
	"github.com/polshe-v/microservices_chat_server/internal/model"
	"github.com/polshe-v/microservices_chat_server/internal/service"
	serviceMocks "github.com/polshe-v/microservices_chat_server/internal/service/mocks"
	desc "github.com/polshe-v/microservices_chat_server/pkg/chat_v1"
)

func TestCreate(t *testing.T) {
	type chatServiceMockFunc func(mc *minimock.Controller) service.ChatService

	type args struct {
		ctx context.Context
		req *desc.CreateRequest
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		id        = int64(1)
		usernames = []string{"name1", "name2", "name3"}

		serviceErr = fmt.Errorf("service error")

		req = &desc.CreateRequest{
			Chat: &desc.Chat{
				Usernames: usernames,
			},
		}

		chat = &model.Chat{
			Usernames: usernames,
		}

		res = &desc.CreateResponse{
			Id: id,
		}
	)

	tests := []struct {
		name            string
		args            args
		want            *desc.CreateResponse
		err             error
		chatServiceMock chatServiceMockFunc
	}{
		{
			name: "service success case",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: res,
			err:  nil,
			chatServiceMock: func(mc *minimock.Controller) service.ChatService {
				mock := serviceMocks.NewChatServiceMock(mc)
				mock.CreateMock.Expect(ctx, chat).Return(id, nil)
				return mock
			},
		},
		{
			name: "service error case",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: nil,
			err:  serviceErr,
			chatServiceMock: func(mc *minimock.Controller) service.ChatService {
				mock := serviceMocks.NewChatServiceMock(mc)
				mock.CreateMock.Expect(ctx, chat).Return(0, serviceErr)
				return mock
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			chatServiceMock := tt.chatServiceMock(mc)
			api := chatAPI.NewImplementation(chatServiceMock)

			res, err := api.Create(tt.args.ctx, tt.args.req)
			require.Equal(t, tt.err, err)
			require.Equal(t, tt.want, res)
		})
	}
}
