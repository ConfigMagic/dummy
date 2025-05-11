package main

import (
	"context"
	"fmt"
	"os"

	"github.com/ConfigMagic/dummy/internal/environment"
	"github.com/spf13/cobra"
)

var upCmd = &cobra.Command{
	Use:   "up",
	Short: "Запустить окружение на основе выбранной конфигурации",
	Long: `Запускает все необходимые сервисы и зависимости для локальной разработки по выбранному конфигу.

Быстрый старт окружения, минимизируя ручные действия и ошибки.`,
	Run: func(cmd *cobra.Command, args []string) {
		mgr := environment.NewDockerComposeManager(configPath)
		err := mgr.Up(context.Background())
		if err != nil {
			fmt.Fprintf(os.Stderr, "Ошибка запуска окружения: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("✅ Окружение запущено")
	},
}

func init() {
	rootCmd.PersistentFlags().StringVar(&configPath, "config", "config.yaml", "Path to the configuration file")
	rootCmd.AddCommand(upCmd)
}
