package main

import (
	"github.com/spf13/cobra"
)

var upCmd = &cobra.Command{
	Use:   "up",
	Short: "Запустить окружение",
	Run: func(cmd *cobra.Command, args []string) {
		runDockerCompose("up", "-d")
	},
}

func init() {
	rootCmd.AddCommand(upCmd)
}
