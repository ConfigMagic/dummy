package composegen

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

// GenerateFromConfig генерирует docker-compose.yaml из пользовательского YAML-конфига (минимально: копирует services)
func GenerateFromConfig(configPath, composePath string) error {
	configBytes, err := os.ReadFile(configPath)
	if err != nil {
		return fmt.Errorf("ошибка чтения конфига: %w", err)
	}
	var config map[string]interface{}
	if err := yaml.Unmarshal(configBytes, &config); err != nil {
		return fmt.Errorf("ошибка парсинга YAML: %w", err)
	}
	compose := map[string]interface{}{
		"version":  "3",
		"services": config["services"],
	}
	composeBytes, err := yaml.Marshal(compose)
	if err != nil {
		return fmt.Errorf("ошибка маршалинга compose: %w", err)
	}
	if err := os.WriteFile(composePath, composeBytes, 0644); err != nil {
		return fmt.Errorf("ошибка сохранения compose: %w", err)
	}
	return nil
}
