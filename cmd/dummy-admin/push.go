package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

var pushCmd = &cobra.Command{
	Use:   "push [config-file]",
	Short: "Опубликовать конфигурацию на сервер (адрес через DUMMY_SERVER_URL)",
	Long: `Загружает или обновляет конфиг окружения на централизованном сервере для всех разработчиков.

Позволяет быстро распространять актуальные настройки и автоматизировать delivery конфигов. Сервер задаётся через переменную окружения DUMMY_SERVER_URL.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		configPath := args[0]

		data, err := os.ReadFile(configPath)
		if err != nil {
			exitWithError(fmt.Errorf("ошибка чтения файла: %v", err))
		}

		var config struct {
			Name string `yaml:"name"`
			Data string `yaml:"data"`
		}
		if err := yaml.Unmarshal(data, &config); err != nil {
			exitWithError(fmt.Errorf("ошибка парсинга YAML: %v", err))
		}

		serverURL := os.Getenv("DUMMY_SERVER_URL")
		if serverURL == "" {
			serverURL = "http://localhost:8080"
		}

		url := serverURL + "/config"
		jsonData := fmt.Sprintf(`{"name":"%s","data":%q}`, config.Name, config.Data)
		resp, err := http.Post(url, "application/json", strings.NewReader(jsonData))
		if err != nil {
			exitWithError(fmt.Errorf("ошибка отправки POST: %v", err))
		}
		defer resp.Body.Close()

		if resp.StatusCode != 200 {
			body, _ := io.ReadAll(resp.Body)
			exitWithError(fmt.Errorf("сервер вернул ошибку: %s", string(body)))
		}

		fmt.Printf("✅ Конфиг '%s' успешно загружен на сервер\n", config.Name)
	},
}

func init() {
	rootCmd.AddCommand(pushCmd)
}
