package root

import (
	"context"
	"fmt"
	"log"

	"github.com/fatih/color"
	"google.golang.org/grpc/metadata"

	desc "github.com/polshe-v/microservices_chat_server/pkg/chat_v1"
)

func createChat(ctx context.Context, address string, certPath string, users []string) error {
	// Get access token from file
	accessToken, err := readToken()
	if err != nil {
		log.Printf("failed to read token: %v", err)
		return err
	}

	client, err := chatServerClient(address, certPath)
	if err != nil {
		return err
	}

	// Get claims from access token
	claims, err := getTokenClaims(accessToken)
	if err != nil {
		return err
	}

	err = isTokenExpired(claims)
	if err != nil {
		return err
	}

	md := metadata.New(map[string]string{"Authorization": "Bearer " + accessToken})
	ctx = metadata.NewOutgoingContext(ctx, md)
	res, err := client.Create(ctx, &desc.CreateRequest{
		Chat: &desc.Chat{
			Usernames: users,
		},
	})
	if err != nil {
		return err
	}

	fmt.Printf("[%s %s]\n", color.CyanString("Created chat with id"), color.YellowString(res.GetId()))
	return nil
}
