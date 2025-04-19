package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var newServerCmd = &cobra.Command{
	Use:   "new_server",
	Short: "Create the dummy's server",
	Long:  `Create the dummy's server for environment management.`,
	Run: func(cmd *cobra.Command, args []string) {
		port, _ := cmd.Flags().GetString("port")
		if port == "" {
			port = "50051"
		}

		err := newServerServer(port)
		if err != nil {
			fmt.Println("Error starting the server:", err)
			os.Exit(1)
		}
		fmt.Println("Server started on port", port)
	},
}

func newServerServer(port string) error {
	cmd := exec.Command("go", "run", "./server/main/main.go", "--port", port)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Start()
}

func init() {
	rootCmd.AddCommand(newServerCmd)

	newServerCmd.Flags().String("port", "", "Port to run the server on (default is 50051)")
}
