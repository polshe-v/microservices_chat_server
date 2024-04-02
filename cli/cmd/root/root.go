package root

import (
	"context"
	"log"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var filename = ".access_token"

// rootCmd represents the base command
var rootCmd = &cobra.Command{
	Use:   "chat-client",
	Short: "Chat client",
}

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login to chat service",
	Run: func(cmd *cobra.Command, _ []string) {
		addr, err := cmd.Flags().GetString("address")
		if err != nil {
			log.Fatalf("failed to get address: %v", err)
		}

		certPath, err := cmd.Flags().GetString("cert")
		if err != nil {
			log.Fatalf("failed to get certificate path: %v", err)
		}

		err = login(context.Background(), addr, certPath)
		if err != nil {
			log.Fatalf("failed to login: %v", err)
		}
	},
}

var logoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "Logout from chat service",
	Run: func(_ *cobra.Command, _ []string) {
		err := logout()
		if err != nil {
			log.Fatalf("failed to logout: %v", err)
		}
	},
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create object",
}

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete object",
}

var connectCmd = &cobra.Command{
	Use:   "connect",
	Short: "Connect to object",
}

var createChatCmd = &cobra.Command{
	Use:   "chat",
	Short: "Create new chat",
	Run: func(cmd *cobra.Command, _ []string) {
		addr, err := cmd.Flags().GetString("address")
		if err != nil {
			log.Fatalf("failed to get address: %v", err)
		}

		certPath, err := cmd.Flags().GetString("cert")
		if err != nil {
			log.Fatalf("failed to get certificate path: %v", err)
		}

		users, err := cmd.Flags().GetString("users")
		if err != nil {
			log.Fatalf("failed to get users: %v", err)
		}

		err = createChat(context.Background(), addr, certPath, strings.Split(users, ","))
		if err != nil {
			log.Fatalf("failed to create chat: %v", err)
		}
	},
}

var deleteChatCmd = &cobra.Command{
	Use:   "chat",
	Short: "Delete chat",
	Run: func(cmd *cobra.Command, _ []string) {
		addr, err := cmd.Flags().GetString("address")
		if err != nil {
			log.Fatalf("failed to get address: %v", err)
		}

		certPath, err := cmd.Flags().GetString("cert")
		if err != nil {
			log.Fatalf("failed to get certificate path: %v", err)
		}

		id, err := cmd.Flags().GetString("id")
		if err != nil {
			log.Fatalf("failed to get chat id: %v", err)
		}

		err = deleteChat(context.Background(), addr, certPath, id)
		if err != nil {
			log.Fatalf("failed to delete chat: %v", err)
		}
	},
}

var connectChatCmd = &cobra.Command{
	Use:   "chat",
	Short: "Connect to chat",
	Run: func(cmd *cobra.Command, _ []string) {
		addr, err := cmd.Flags().GetString("address")
		if err != nil {
			log.Fatalf("failed to get address: %v", err)
		}

		certPath, err := cmd.Flags().GetString("cert")
		if err != nil {
			log.Fatalf("failed to get certificate path: %v", err)
		}

		id, err := cmd.Flags().GetString("id")
		if err != nil {
			log.Fatalf("failed to get chat id: %v", err)
		}

		err = connectChat(context.Background(), addr, certPath, id)
		if err != nil {
			log.Fatalf("failed to connect to chat: %v", err)
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(loginCmd)
	rootCmd.AddCommand(logoutCmd)
	rootCmd.AddCommand(createCmd)
	rootCmd.AddCommand(deleteCmd)
	rootCmd.AddCommand(connectCmd)

	createCmd.AddCommand(createChatCmd)
	deleteCmd.AddCommand(deleteChatCmd)
	connectCmd.AddCommand(connectChatCmd)

	// Add flags to commands
	loginCmd.Flags().StringP("address", "a", "", "`IP:port` of authentication service")
	loginCmd.Flags().StringP("cert", "c", "", "path to authentication service certificate")
	createChatCmd.Flags().StringP("users", "u", "", "chat participants")
	createChatCmd.Flags().StringP("address", "a", "", "`IP:port` of chat service")
	createChatCmd.Flags().StringP("cert", "c", "", "path to chat service certificate")
	deleteChatCmd.Flags().StringP("id", "n", "", "ID of the chat to delete")
	deleteChatCmd.Flags().StringP("address", "a", "", "`IP:port` of chat service")
	deleteChatCmd.Flags().StringP("cert", "c", "", "path to chat service certificate")
	connectChatCmd.Flags().StringP("id", "n", "", "ID of the chat to connect")
	connectChatCmd.Flags().StringP("address", "a", "", "`IP:port` of chat service")
	connectChatCmd.Flags().StringP("cert", "c", "", "path to chat service certificate")

	// Mark required flags in commands
	err := loginCmd.MarkFlagRequired("address")
	if err != nil {
		log.Fatalf("failed to mark address flag as required: %v", err)
	}

	err = loginCmd.MarkFlagRequired("cert")
	if err != nil {
		log.Fatalf("failed to mark cert flag as required: %v", err)
	}

	err = createChatCmd.MarkFlagRequired("address")
	if err != nil {
		log.Fatalf("failed to mark id flag as required: %v", err)
	}

	err = createChatCmd.MarkFlagRequired("cert")
	if err != nil {
		log.Fatalf("failed to mark cert flag as required: %v", err)
	}

	err = deleteChatCmd.MarkFlagRequired("id")
	if err != nil {
		log.Fatalf("failed to mark id flag as required: %v", err)
	}

	err = deleteChatCmd.MarkFlagRequired("address")
	if err != nil {
		log.Fatalf("failed to mark address flag as required: %v", err)
	}

	err = deleteChatCmd.MarkFlagRequired("cert")
	if err != nil {
		log.Fatalf("failed to mark cert flag as required: %v", err)
	}

	err = connectChatCmd.MarkFlagRequired("id")
	if err != nil {
		log.Fatalf("failed to mark id flag as required: %v", err)
	}

	err = connectChatCmd.MarkFlagRequired("address")
	if err != nil {
		log.Fatalf("failed to mark address flag as required: %v", err)
	}

	err = connectChatCmd.MarkFlagRequired("cert")
	if err != nil {
		log.Fatalf("failed to mark cert flag as required: %v", err)
	}
}
