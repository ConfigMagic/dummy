package main

import "github.com/spf13/cobra"

var downCmd = &cobra.Command{
	Use:   "down",
	Short: "Stop the environment",
	Long:  "... some long description ...",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Println("Stopping the environment")
	},
}

func init() {
	rootCmd.AddCommand(downCmd)
}
