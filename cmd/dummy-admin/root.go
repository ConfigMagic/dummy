package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "dummy-admin",
	Short: "Быстрое администрирование и централизованное управление окружениями (one-click setup)",
	Long: `dummy-admin — инструмент для администраторов и DevOps-инженеров для централизованного управления конфигурациями, пользователями и сервисами.

Основные возможности:
- Публикация и обновление конфигов для разработчиков
- Управление пользователями и доступами
- Интеграция с сервером и централизованным хранилищем
- Генерация скриптов автодополнения для shell

Упрощает поддержку, ускоряет распространение изменений и снижает количество ручных операций при администрировании окружений.`,
}

var completionCmd = &cobra.Command{
	Use:   "completion",
	Short: "Сгенерировать скрипт автодополнения для shell (bash/zsh)",
	Long: `Автодополнение команд dummy-admin для удобства администрирования.

Supported shells:
  - bash
  - zsh

Usage:
  dummy-admin completion bash > completions/dummy-admin.bash
  dummy-admin completion zsh > completions/_dummy-admin`,
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
