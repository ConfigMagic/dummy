package main

import (
	"context"
	"fmt"
	"os"

	"github.com/ConfigMagic/dummy/internal/environment"
	"github.com/spf13/cobra"
)

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Показать текущее состояние окружения",
	Run: func(cmd *cobra.Command, args []string) {
		mgr := environment.NewDockerComposeManager(configPath)
		status, err := mgr.Status(context.Background())
		if err != nil {
			fmt.Fprintf(os.Stderr, "Ошибка получения статуса: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("Статус окружения:")
		fmt.Println(status)
	},
}

func init() {
	rootCmd.AddCommand(statusCmd)
}
