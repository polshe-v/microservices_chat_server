package auth

import (
	"context"

	desc "github.com/polshe-v/microservices_auth/internal/pkg/access_v1"
	"github.com/polshe-v/microservices_chat_server/internal/client/rpc"
)

type Client struct {
	client desc.AccessV1Client
}

var _ rpc.AuthClient = (*Client)(nil)

// NewGrpcConfig creates new object of GrpcConfig interface.
func NewAuthClient(client desc.AccessV1Client) *Client {
	return &Client{
		client: client,
	}
}

func (c *Client) Check(ctx context.Context, endpoint string) error {
	_, err := c.client.Check(ctx, &desc.CheckRequest{
		EndpointAddress: endpoint,
	})
	return err
}
