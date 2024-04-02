package root

import (
	"bufio"
	"context"
	"fmt"
	"os"

	"github.com/fatih/color"
	"golang.org/x/term"

	desc "github.com/polshe-v/microservices_auth/pkg/auth_v1"
)

func login(ctx context.Context, address string, certPath string) error {
	client, err := authClient(address, certPath)
	if err != nil {
		return err
	}

	var username string

	fmt.Print(color.CyanString("Username: "))
	fmt.Scan(&username)

	fmt.Print(color.CyanString("Password (no echo): "))
	password, err := term.ReadPassword(0)
	if err != nil {
		return err
	}

	// Get refresh token
	res, err := client.Login(ctx, &desc.LoginRequest{
		Creds: &desc.Creds{
			Username: username,
			Password: string(password),
		},
	})
	if err != nil {
		return err
	}

	// Get access token
	resAccessToken, err := client.GetAccessToken(ctx, &desc.GetAccessTokenRequest{
		RefreshToken: res.GetRefreshToken(),
	})
	if err != nil {
		return err
	}

	// Save access token in file for later operations
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	w := bufio.NewWriter(f)
	_, err = w.WriteString(resAccessToken.GetAccessToken())
	if err != nil {
		return err
	}
	w.Flush()

	fmt.Println(color.GreenString("\n\n[Successfully logged in]\n"))
	return nil
}
