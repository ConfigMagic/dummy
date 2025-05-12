package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

func indentYaml(data string) string {
	if data == "" {
		return ""
	}
	lines := strings.Split(data, "\n")
	for i, line := range lines {
		lines[i] = "  " + line
	}
	return strings.Join(lines, "\n")
}

var syncCmd = &cobra.Command{
	Use:   "sync [название_конфига]",
	Short: "Синхронизировать конфиг с сервером (адрес через DUMMY_SERVER_URL)",
	Long: `Получить или обновить локальный конфиг из централизованного хранилища.

Позволяет держать конфиги в актуальном состоянии для всех разработчиков.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		configName := args[0]

		serverURL := os.Getenv("DUMMY_SERVER_URL")
		if serverURL == "" {
			serverURL = "http://localhost:8080"
		}
		url := fmt.Sprintf("%s/config/%s", serverURL, configName)

		resp, err := http.Get(url)
		if err != nil {
			exitWithError(fmt.Errorf("ошибка запроса к серверу: %v", err))
		}
		defer resp.Body.Close()

		if resp.StatusCode != 200 {
			body, _ := io.ReadAll(resp.Body)
			exitWithError(fmt.Errorf("сервер вернул ошибку: %s", string(body)))
		}

		var result struct {
			Data string `json:"data"`
		}
		if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
			exitWithError(fmt.Errorf("ошибка декодирования ответа: %v", err))
		}

		configYaml := fmt.Sprintf("data: |\n%s\n", indentYaml(result.Data))
		homeDir, err := os.UserHomeDir()
		if err != nil {
			exitWithError(fmt.Errorf("не удалось определить домашний каталог: %v", err))
		}
		dummyDir := filepath.Join(homeDir, ".dummy")
		if err := os.MkdirAll(dummyDir, 0755); err != nil {
			exitWithError(fmt.Errorf("не удалось создать директорию ~/.dummy: %v", err))
		}
		filename := filepath.Join(dummyDir, fmt.Sprintf("%s.yaml", configName))
		if err := os.WriteFile(filename, []byte(configYaml), 0644); err != nil {
			exitWithError(fmt.Errorf("ошибка сохранения конфига: %v", err))
		}

		fmt.Printf("✅ Конфиг '%s' загружен\n", configName)
		fmt.Printf("Конфиг сохранён в файл: %s\n", filename)
	},
}

func init() {
	rootCmd.AddCommand(syncCmd)
}
