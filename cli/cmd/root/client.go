package root

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	descAuth "github.com/polshe-v/microservices_auth/pkg/auth_v1"
	descChat "github.com/polshe-v/microservices_chat_server/pkg/chat_v1"
)

func authClient(address string, certPath string) (descAuth.AuthV1Client, error) {
	creds, err := credentials.NewClientTLSFromFile(certPath, "")
	if err != nil {
		log.Fatalf("failed to get credentials of authentication service: %v", err)
	}

	conn, err := grpc.Dial(address,
		grpc.WithTransportCredentials(creds),
	)
	if err != nil {
		log.Fatalf("failed to connect to auth server: %v", err)
	}
	//defer conn.Close()

	return descAuth.NewAuthV1Client(conn), nil
}

func chatServerClient(address string, certPath string) (descChat.ChatV1Client, error) {
	creds, err := credentials.NewClientTLSFromFile(certPath, "")
	if err != nil {
		log.Fatalf("failed to get credentials of chat service: %v", err)
	}

	conn, err := grpc.Dial(address,
		grpc.WithTransportCredentials(creds),
	)
	if err != nil {
		log.Fatalf("failed to connect to chat server: %v", err)
	}
	//defer conn.Close()
	return descChat.NewChatV1Client(conn), nil
}
