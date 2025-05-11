package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "dummy",
	Short: "a one-click setup environment tool for developers",
	Long:  "... some long description ...",
}

var completionCmd = &cobra.Command{
	Use:   "completion",
	Short: "Generate shell completion script",
	Run: func(cmd *cobra.Command, args []string) {
		rootCmd.GenBashCompletion(os.Stdout) // можно сделать switch для zsh/fish
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		exitWithError(err)
	}
}

func exitWithError(err error) {
	fmt.Fprintf(os.Stderr, "Error: %v\n", err)
	os.Exit(1)
}

func init() {
	rootCmd.AddCommand(completionCmd)
	// Добавляем новые команды
	rootCmd.AddCommand(reloadCmd)
	rootCmd.AddCommand(statusCmd)
}
