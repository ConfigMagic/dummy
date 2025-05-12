package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var configPath string

var rootCmd = &cobra.Command{
	Use:   "dummy",
	Short: "Быстрый запуск и управление окружением для разработчиков (one-click setup)",
	Long: `dummy — инструмент для быстрого старта, управления и обновления локального окружения разработчика.

Основные возможности:
- Запуск и остановка сервисов по конфигу
- Динамическое обновление окружения
- Получение статуса и логов
- Синхронизация конфигов с сервером
- Генерация docker-compose.yaml

Упрощает онбординг, снижает количество ручных действий и ошибок при работе с микросервисами и зависимостями.`,
}

var completionCmd = &cobra.Command{
	Use:   "completion",
	Short: "Сгенерировать скрипт автодополнения для shell (bash/zsh)",
	Long:  "Удобное автодополнение команд dummy в терминале. Пример: dummy completion bash > completions/dummy.bash",
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
	rootCmd.PersistentFlags().StringVarP(&configPath, "config", "c", "config.yaml", "Путь к конфигурационному файлу")
}
