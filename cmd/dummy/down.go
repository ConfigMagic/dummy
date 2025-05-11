package main

import (
	"context"
	"fmt"
	"os"

	"github.com/ConfigMagic/dummy/internal/environment"
	"github.com/spf13/cobra"
)

var downCmd = &cobra.Command{
	Use:   "down",
	Short: "Остановить окружение и освободить ресурсы",
	Long: `Полностью выключает все сервисы, поднятые через dummy, и освобождает ресурсы.

Удобно для завершения работы или очистки среды.`,
	Run: func(cmd *cobra.Command, args []string) {
		mgr := environment.NewDockerComposeManager(configPath)
		err := mgr.Down(context.Background())
		if err != nil {
			fmt.Fprintf(os.Stderr, "Ошибка остановки окружения: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("✅ Окружение остановлено")
	},
}

func init() {
	rootCmd.AddCommand(downCmd)
	downCmd.PersistentFlags().StringVar(&configPath, "config", "config.yaml", "Путь к конфигурационному файлу")
}
