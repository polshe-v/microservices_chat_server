package root

import (
	"log"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command
var rootCmd = &cobra.Command{
	Use:   "chat-client",
	Short: "Chat client",
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create something",
}

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete something",
}

var createUserCmd = &cobra.Command{
	Use:   "user",
	Short: "Create new user",
	Run: func(cmd *cobra.Command, _ []string) {
		usernamesStr, err := cmd.Flags().GetString("username")
		if err != nil {
			log.Fatalf("failed to get usernames: %v", err)
		}

		log.Printf("user %s created\n", usernamesStr)
	},
}

var deleteUserCmd = &cobra.Command{
	Use:   "user",
	Short: "Delete user",
	Run: func(cmd *cobra.Command, _ []string) {
		usernamesStr, err := cmd.Flags().GetString("username")
		if err != nil {
			log.Fatalf("failed to get usernames: %v", err)
		}

		log.Printf("user %s deleted\n", usernamesStr)
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
	rootCmd.AddCommand(createCmd)
	rootCmd.AddCommand(deleteCmd)

	createCmd.AddCommand(createUserCmd)
	deleteCmd.AddCommand(deleteUserCmd)

	createUserCmd.Flags().StringP("username", "u", "", "Name of the user")
	err := createUserCmd.MarkFlagRequired("username")
	if err != nil {
		log.Fatalf("failed to mark username flag as required: %v", err)
	}

	deleteUserCmd.Flags().StringP("username", "u", "", "Name of the user")
	err = deleteUserCmd.MarkFlagRequired("username")
	if err != nil {
		log.Fatalf("failed to mark username flag as required: %v", err)
	}
}
