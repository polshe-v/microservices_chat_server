package auth

import (
	"context"

	desc "github.com/polshe-v/microservices_auth/internal/pkg/access_v1"
	"github.com/polshe-v/microservices_chat_server/internal/client/rpc"
)

// Client contains client connection with authentication service.
type Client struct {
	Client desc.AccessV1Client
}

var _ rpc.AuthClient = (*Client)(nil)

// NewAuthClient creates new Client object.
func NewAuthClient(client desc.AccessV1Client) *Client {
	return &Client{
		Client: client,
	}
}

// Check calls authentication service method for authorization.
func (c *Client) Check(ctx context.Context, endpoint string) error {
	_, err := c.Client.Check(ctx, &desc.CheckRequest{
		EndpointAddress: endpoint,
	})
	return err
}
