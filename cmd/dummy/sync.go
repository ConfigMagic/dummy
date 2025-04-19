package main

import "github.com/spf13/cobra"

var syncCmd = &cobra.Command{
	Use:   "sync",
	Short: "Synchronize the configurations",
	Long:  "... some long description ...",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Println("Synchronizing the environment")
	},
}

func init() {
	rootCmd.AddCommand(syncCmd)
}
