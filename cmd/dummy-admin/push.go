package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

var pushCmd = &cobra.Command{
	Use:   "push",
	Short: "Publish configuration to the server",
	Long:  "... some long description ...",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Configuration pushed to the server")
	},
}

func init() {
	rootCmd.AddCommand(pushCmd)
}
