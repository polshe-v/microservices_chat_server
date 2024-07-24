package root

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	descAuth "github.com/polshe-v/microservices_auth/pkg/auth_v1"
	descChat "github.com/polshe-v/microservices_chat_server/pkg/chat_v1"
	"github.com/polshe-v/microservices_common/pkg/closer"
)

func authClient(address string, certPath string) (descAuth.AuthV1Client, error) {
	creds, err := credentials.NewClientTLSFromFile(certPath, "")
	if err != nil {
		return nil, err
	}

	conn, err := grpc.NewClient(address,
		grpc.WithTransportCredentials(creds),
	)
	if err != nil {
		return nil, err
	}
	closer.Add(conn.Close)

	return descAuth.NewAuthV1Client(conn), nil
}

func chatServerClient(address string, certPath string) (descChat.ChatV1Client, error) {
	creds, err := credentials.NewClientTLSFromFile(certPath, "")
	if err != nil {
		return nil, err
	}

	conn, err := grpc.NewClient(address,
		grpc.WithTransportCredentials(creds),
	)
	if err != nil {
		return nil, err
	}
	closer.Add(conn.Close)

	return descChat.NewChatV1Client(conn), nil
}
