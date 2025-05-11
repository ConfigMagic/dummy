package main

import (
	"context"
	"fmt"
	"io/ioutil"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

var syncCmd = &cobra.Command{
	Use:   "sync [название_конфига]",
	Short: "Синхронизировать конфигурацию с сервером",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		configName := args[0]
		client, err := NewGRPCClient("localhost:50051")
		if err != nil {
			exitWithError(err)
		}
		defer client.Close()

		ctx := context.Background()
		config, err := client.GetConfig(ctx, configName)
		if err != nil {
			exitWithError(fmt.Errorf("ошибка получения конфига: %v", err))
		}

		// Сохранить конфиг в YAML
		configBytes, _ := yaml.Marshal(config)
		filename := fmt.Sprintf("%s.yaml", configName)
		if err := ioutil.WriteFile(filename, configBytes, 0644); err != nil {
			exitWithError(fmt.Errorf("ошибка сохранения конфига: %v", err))
		}

		fmt.Printf("✅ Конфиг '%s' загружен\n", config.Name)
		fmt.Printf("Конфиг сохранён в файл: %s\n", filename)
	},
}

func init() {
	rootCmd.AddCommand(syncCmd)
}
