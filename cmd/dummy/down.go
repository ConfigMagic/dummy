package main

import (
	"github.com/spf13/cobra"
)

var downCmd = &cobra.Command{
	Use:   "down",
	Short: "Остановить окружение",
	Run: func(cmd *cobra.Command, args []string) {
		runDockerCompose("down")
	},
}

func init() {
	rootCmd.AddCommand(downCmd)
}
