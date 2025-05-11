package main

import (
	"context"
	"fmt"
	"os"

	"github.com/ConfigMagic/dummy/internal/environment"
	"github.com/spf13/cobra"
)

var reloadCmd = &cobra.Command{
	Use:   "reload",
	Short: "Динамически обновить окружение (без полного пересоздания)",
	Run: func(cmd *cobra.Command, args []string) {
		mgr := environment.NewDockerComposeManager(configPath)
		err := mgr.Reload(context.Background())
		if err != nil {
			fmt.Fprintf(os.Stderr, "Ошибка обновления окружения: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("✅ Окружение обновлено (reload)")
	},
}

func init() {
	rootCmd.AddCommand(reloadCmd)
}
