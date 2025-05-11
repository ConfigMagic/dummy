package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var newServerCmd = &cobra.Command{
	Use:   "new_server",
	Short: "Запустить сервер dummy для централизованного управления окружениями",
	Long: `Запускает серверную часть dummy для централизованного хранения, выдачи и управления конфигурациями окружений.

Позволяет администраторам и DevOps быстро развернуть backend для работы с конфигами, пользователями и сервисами. Поддерживает настройку порта запуска.`,
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
	},
}

func newServerServer(port string) error {
	// Сначала проверим, что все зависимости установлены
	cmd := exec.Command("go", "mod", "tidy")
	cmd.Dir = "./server"
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to update dependencies: %v", err)
	}

	// Теперь запустим сервер
	cmd = exec.Command("go", "run", "cmd/app/main.go")
	cmd.Dir = "./server"
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Start(); err != nil {
		return fmt.Errorf("failed to start server: %v", err)
	}

	fmt.Println("Server started on port", port)
	return nil
}

func init() {
	rootCmd.AddCommand(newServerCmd)

	newServerCmd.Flags().String("port", "", "Port to run the server on (default is 50051)")
}
