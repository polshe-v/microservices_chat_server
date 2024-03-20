package interceptor

import (
	"context"

	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	"github.com/polshe-v/microservices_chat_server/internal/client/rpc/auth"
)

// Client contains client connection with authentication service.
type Client struct {
	Client *auth.Client
}

// PolicyInterceptor is used for authorization.
func (c *Client) PolicyInterceptor(ctx context.Context, req interface{}, _ *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	endpoint, ok := req.(string)
	if !ok {
		return nil, errors.New("invalid endpoint type")
	}

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errors.New("metadata is not provided")
	}

	err := c.Client.Check(metadata.NewOutgoingContext(ctx, md), endpoint)
	if err != nil {
		return nil, err
	}

	return handler(ctx, req)
}
