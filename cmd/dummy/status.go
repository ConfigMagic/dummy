package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Показать текущее состояние окружения",
	Long: `Показывает, что сейчас запущено и в каком статусе находятся сервисы.

Удобно для быстрой диагностики и контроля состояния среды.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Fprintln(os.Stderr, "Команда reload пока не поддерживается в универсальном runner.")
		os.Exit(1)
	},
}

func init() {
	rootCmd.AddCommand(statusCmd)
}
