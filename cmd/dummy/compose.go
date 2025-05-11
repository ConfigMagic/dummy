package main

import (
	"fmt"
	"io/ioutil"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

var composeCmd = &cobra.Command{
	Use:   "generate [название_конфига]",
	Short: "Сгенерировать docker-compose.yaml из выбранного конфига",
	Long: `Автоматически собирает docker-compose файл для ручного запуска или отладки сервисов.

Удобно для интеграции с существующими инструментами и CI/CD.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		configName := args[0]
		configBytes, err := ioutil.ReadFile(fmt.Sprintf("%s.yaml", configName))
		if err != nil {
			exitWithError(fmt.Errorf("ошибка чтения конфига: %v", err))
		}

		var config map[string]interface{}
		if err := yaml.Unmarshal(configBytes, &config); err != nil {
			exitWithError(fmt.Errorf("ошибка парсинга YAML: %v", err))
		}

		// Пример простой генерации docker-compose
		compose := map[string]interface{}{
			"version":  "3",
			"services": config["services"],
		}

		composeBytes, _ := yaml.Marshal(compose)
		composeFilename := "docker-compose.yaml"
		if err := ioutil.WriteFile(composeFilename, composeBytes, 0644); err != nil {
			exitWithError(fmt.Errorf("ошибка сохранения docker-compose.yaml: %v", err))
		}

		fmt.Printf("✅ docker-compose.yaml создан в %s\n", composeFilename)
	},
}

func init() {
	rootCmd.AddCommand(composeCmd)
}
