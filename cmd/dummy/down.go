package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/ConfigMagic/dummy/internal/runner"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

var downCmd = &cobra.Command{
	Use:   "down",
	Short: "Остановить окружение и освободить ресурсы",
	Long: `Полностью выключает все сервисы, поднятые через dummy, и освобождает ресурсы.

Удобно для завершения работы или очистки среды.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Определяем имя окружения (например, server)
		configName := configPath
		if filepath.Ext(configName) == ".yaml" {
			configName = configName[:len(configName)-5]
		}
		configDir := filepath.Dir(configPath)
		baseName := filepath.Base(configName)
		envDir := filepath.Join(configDir, baseName)
		runnerPath := filepath.Join(envDir, "runner.yaml")

		// Проверяем runner.yaml по основному пути, иначе ищем в ~/.dummy
		if _, err := os.Stat(runnerPath); os.IsNotExist(err) {
			homeDir, _ := os.UserHomeDir()
			altEnvDir := filepath.Join(homeDir, ".dummy", baseName)
			altRunnerPath := filepath.Join(altEnvDir, "runner.yaml")
			if _, err := os.Stat(altRunnerPath); err == nil {
				runnerPath = altRunnerPath
				envDir = altEnvDir
			} else {
				fmt.Fprintf(os.Stderr, "runner.yaml не найден ни по основному пути, ни в ~/.dummy\n")
				os.Exit(1)
			}
		}

		cfg, err := runner.LoadRunnerConfig(runnerPath)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Ошибка чтения runner.yaml: %v\n", err)
			os.Exit(1)
		}
		// Читаем env из user yaml
		userConfigBytes, err := os.ReadFile(configPath)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Ошибка чтения user config: %v\n", err)
			os.Exit(1)
		}
		var userConfig map[string]interface{}
		yaml.Unmarshal(userConfigBytes, &userConfig)
		env := map[string]string{}
		if envSection, ok := userConfig["env"].(map[interface{}]interface{}); ok {
			for k, v := range envSection {
				ks, vs := fmt.Sprintf("%v", k), fmt.Sprintf("%v", v)
				env[ks] = vs
			}
		}
		// Запускаем down_command
		if err := runner.RunDownCommand(cfg, env, envDir); err != nil {
			fmt.Fprintf(os.Stderr, "Ошибка остановки runner: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("✅ Окружение остановлено через универсальный runner")
	},
}

func init() {
	rootCmd.AddCommand(downCmd)
	downCmd.PersistentFlags().StringVar(&configPath, "config", "config.yaml", "Путь к конфигурационному файлу")
}
