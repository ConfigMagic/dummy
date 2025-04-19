package main

import "github.com/spf13/cobra"

var upCmd = &cobra.Command{
	Use:   "up",
	Short: "Start the environment",
	Long:  "... some long description ...",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Println("Starting the environment")
	},
}

func init() {
	rootCmd.AddCommand(upCmd)
}
