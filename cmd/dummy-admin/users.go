package main

import "github.com/spf13/cobra"

var usersCmd = &cobra.Command{
	Use:   "users",
	Short: "Manage users",
	Long:  "... some long description ...",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Println("Manage users")
	},
}

func init() {
	rootCmd.AddCommand(usersCmd)
}
