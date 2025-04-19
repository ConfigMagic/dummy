package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "dummy-admin",
	Short: "a one-click setup environment tool for administrators",
	Long:  "... some long description ...",
}

var completionCmd = &cobra.Command{
	Use:   "completion",
	Short: "Generate shell completion script",
	Long: `Generate shell autocompletion script.

Supported shells:
  - bash
  - zsh

Usage:
  dummy completion bash > completions/dummy.bash
  dummy completion zsh > completions/_dummy`,
	Args: cobra.ExactArgs(1), // только один аргумент: shell
	RunE: func(cmd *cobra.Command, args []string) error {
		shell := args[0]

		switch shell {
		case "bash":
			return rootCmd.GenBashCompletion(os.Stdout)
		case "zsh":
			return rootCmd.GenZshCompletion(os.Stdout)
		default:
			return fmt.Errorf("unsupported shell: %s", shell)
		}
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
}
