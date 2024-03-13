package interceptor

import (
	"context"

	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"

	desc "github.com/polshe-v/microservices_auth/internal/pkg/access_v1"
)

type AuthServiceParams struct {
	AuthAddress  string
	AuthCertPath string
}

var AuthService *AuthServiceParams

func PolicyInterceptor(ctx context.Context, req interface{}, _ *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	endpoint, ok := req.(string)
	if !ok {
		return nil, errors.New("invalid endpoint type provided")
	}

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errors.New("metadata is not provided")
	}

	creds, err := credentials.NewClientTLSFromFile(AuthService.AuthCertPath, "")
	if err != nil {
		return nil, err
	}

	conn, err := grpc.Dial(
		AuthService.AuthAddress,
		grpc.WithTransportCredentials(creds),
	)
	if err != nil {
		return nil, errors.New("failed to connect to authentication service")
	}

	cl := desc.NewAccessV1Client(conn)

	_, err = cl.Check(metadata.NewOutgoingContext(ctx, md), &desc.CheckRequest{
		EndpointAddress: endpoint,
	})

	if err != nil {
		return nil, err
	}

	return handler(ctx, req)
}
