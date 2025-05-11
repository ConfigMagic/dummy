package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var reloadCmd = &cobra.Command{
	Use:   "reload",
	Short: "Динамически обновить окружение (без полного пересоздания)",
	Long:  "Пока не поддерживается в универсальном runner. Используйте dummy down/up.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Fprintln(os.Stderr, "Команда reload пока не поддерживается в универсальном runner. Используйте dummy down/up.")
		os.Exit(1)
	},
}

func init() {
	rootCmd.AddCommand(reloadCmd)
}
