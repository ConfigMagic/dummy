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
	Short: "Остановить окружение",
	Run: func(cmd *cobra.Command, args []string) {
		mgr := environment.NewDockerComposeManager("config.yaml") // TODO: путь к конфигу сделать параметром
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
}
