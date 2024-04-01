package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/fatih/color"
	"golang.org/x/term"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/types/known/timestamppb"

	descAuth "github.com/polshe-v/microservices_auth/pkg/auth_v1"
	desc "github.com/polshe-v/microservices_chat_server/pkg/chat_v1"
	//"github.com/polshe-v/microservices_chat_server/cli/cmd/root"
)

const (
	addressAuth = "localhost:50000"
	address     = "localhost:50001"
)

func main() {
	//root.Execute()
	var (
		chatID   string
		username string
	)
	ctx := context.Background()

	fmt.Println(color.CyanString("[Login]"))

	fmt.Print("Username: ")
	fmt.Scan(&username)

	fmt.Print("Password (no echo): ")
	password, err := term.ReadPassword(0)
	if err != nil {
		log.Fatalf("failed to read password: %v", err)
	}

	credsAuth, err := credentials.NewClientTLSFromFile("tls/auth.pem", "")
	if err != nil {
		log.Fatalf("failed to get credentials of authentication service: %v", err)
	}

	connAuth, err := grpc.Dial(addressAuth,
		grpc.WithTransportCredentials(credsAuth),
	)
	if err != nil {
		log.Fatalf("failed to connect to auth server: %v", err)
	}
	defer connAuth.Close()

	clientAuth := descAuth.NewAuthV1Client(connAuth)
	accessToken, err := login(ctx, clientAuth, username, string(password))
	if err != nil {
		log.Fatalf("failed to login: %v", err)
	}

	fmt.Println(color.CyanString("\n[Connect]"))

	fmt.Print("ChatID: ")
	fmt.Scan(&chatID)

	fmt.Println(color.CyanString("[Chat]"))

	creds, err := credentials.NewClientTLSFromFile("tls/chat.pem", "")
	if err != nil {
		log.Fatalf("failed to get credentials of chat service: %v", err)
	}

	conn, err := grpc.Dial(address,
		grpc.WithTransportCredentials(creds),
	)
	if err != nil {
		log.Fatalf("failed to connect to chat server: %v", err)
	}
	defer conn.Close()

	client := desc.NewChatV1Client(conn)
	md := metadata.New(map[string]string{"Authorization": "Bearer " + accessToken})
	err = connectChat(metadata.NewOutgoingContext(ctx, md), client, chatID, username)
	if err != nil {
		log.Fatalf("failed to connect to chatID: %v", err)
	}
}

func login(ctx context.Context, client descAuth.AuthV1Client, username string, password string) (string, error) {
	resLogin, err := client.Login(ctx, &descAuth.LoginRequest{
		Creds: &descAuth.Creds{
			Username: username,
			Password: password,
		},
	})
	if err != nil {
		log.Print("Login error")
		return "", err
	}

	resAccessToken, err := client.GetAccessToken(ctx, &descAuth.GetAccessTokenRequest{
		RefreshToken: resLogin.GetRefreshToken(),
	})
	if err != nil {
		log.Print("Get accessToken error")
		return "", err
	}

	return resAccessToken.GetAccessToken(), nil
}

func connectChat(ctx context.Context, client desc.ChatV1Client, chatID string, username string) error {
	stream, err := client.Connect(ctx, &desc.ConnectRequest{
		ChatId:   chatID,
		Username: username,
	})
	if err != nil {
		log.Print("Connect error")
		return err
	}

	go func() {
		for {
			message, errRecv := stream.Recv()
			if errRecv == io.EOF {
				log.Print("Recv error")
				return
			}
			if errRecv != nil {
				log.Printf("failed to receive message from stream: %v", err)
				return
			}

			fmt.Printf("[%v] - [from: %s]: %s",
				color.YellowString(message.GetTimestamp().AsTime().Format(time.RFC3339)),
				color.GreenString(message.GetFrom()),
				message.GetText(),
			)
		}
	}()

	for {
		reader := bufio.NewReader(os.Stdin)
		text, err := reader.ReadString('\n')
		if err != nil {
			return err
		}

		// Input will be replaced with formatted message later
		fmt.Printf("\033[1A")
		_, err = client.SendMessage(ctx, &desc.SendMessageRequest{
			ChatId: chatID,
			Message: &desc.Message{
				From:      username,
				Text:      text,
				Timestamp: timestamppb.Now(),
			},
		})
		if err != nil {
			log.Printf("failed to send message: %v", err)
		}
	}
}
