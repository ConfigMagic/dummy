package main

import (
	"context"
	"fmt"
	"io/ioutil"

	pb "github.com/ConfigMagic/dummy/server/pb"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

var pushCmd = &cobra.Command{
	Use:   "push [config-file]",
	Short: "Publish configuration to the server",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		configPath := args[0]

		// Чтение файла конфигурации
		data, err := ioutil.ReadFile(configPath)
		if err != nil {
			exitWithError(fmt.Errorf("ошибка чтения файла: %v", err))
		}

		// Парсинг YAML
		var config pb.EnvConfig
		if err := yaml.Unmarshal(data, &config); err != nil {
			exitWithError(fmt.Errorf("ошибка парсинга YAML: %v", err))
		}

		// Подключение к серверу
		client, err := NewGRPCAdminClient("localhost:50051")
		if err != nil {
			exitWithError(err)
		}
		defer client.Close()

		// Отправка конфига на сервер
		ctx := context.Background()
		_, err = client.ApplyConfig(ctx, &config)
		if err != nil {
			exitWithError(fmt.Errorf("ошибка применения конфига: %v", err))
		}

		fmt.Printf("✅ Конфиг '%s' успешно загружен на сервер\n", config.Name)
	},
}

func init() {
	rootCmd.AddCommand(pushCmd)
}
