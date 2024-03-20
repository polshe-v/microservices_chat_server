package auth

import (
	"context"

	desc "github.com/polshe-v/microservices_auth/internal/pkg/access_v1"
	"github.com/polshe-v/microservices_chat_server/internal/client/rpc"
)

type client struct {
	client desc.AccessV1Client
}

var _ rpc.AuthClient = (*client)(nil)

// NewAuthClient creates new AuthClient object.
func NewAuthClient(client desc.AccessV1Client) rpc.AuthClient {
	return &client{
		client: client,
	}
}

// Check calls authentication service method for authorization.
func (c *client) Check(ctx context.Context, endpoint string) error {
	_, err := c.client.Check(ctx, &desc.CheckRequest{
		EndpointAddress: endpoint,
	})
	return err
}
