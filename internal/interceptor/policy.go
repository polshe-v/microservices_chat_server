package interceptor

import (
	"context"

	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	"github.com/polshe-v/microservices_auth/internal/client/rpc/auth"
	desc "github.com/polshe-v/microservices_auth/internal/pkg/access_v1"
)

type Client struct {
	client *auth.Client
}

func (c *Client) PolicyInterceptor(ctx context.Context, req interface{}, _ *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	endpoint, ok := req.(string)
	if !ok {
		return nil, errors.New("invalid endpoint type")
	}

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errors.New("metadata is not provided")
	}

	err := c.client.Check(metadata.NewOutgoingContext(ctx, md), endpoint)
	if err != nil {
		return nil, err
	}

	return handler(ctx, req)
}
