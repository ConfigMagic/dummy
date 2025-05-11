package environment

import (
	"context"
	"fmt"
	"os/exec"

	"github.com/ConfigMagic/dummy/internal/composegen"
)

// EnvironmentManager описывает интерфейс для управления окружением (создание, уничтожение, обновление, статус)
type EnvironmentManager interface {
	Up(ctx context.Context) error
	Down(ctx context.Context) error
	Reload(ctx context.Context) error
	Status(ctx context.Context) (string, error)
}

// DockerComposeManager — базовая реализация через docker compose
// В будущем можно добавить HelmManager, K8sManager и т.д.
type DockerComposeManager struct {
	ConfigPath string
}

func NewDockerComposeManager(configPath string) *DockerComposeManager {
	return &DockerComposeManager{ConfigPath: configPath}
}

func (d *DockerComposeManager) Up(ctx context.Context) error {
	// 1. Генерируем docker-compose.yaml из конфига
	composePath := "docker-compose.yaml"
	err := composegen.GenerateFromConfig(d.ConfigPath, composePath)
	if err != nil {
		return fmt.Errorf("ошибка генерации compose: %w", err)
	}
	// 2. Запускаем docker compose up -d
	cmd := exec.CommandContext(ctx, "docker", "compose", "-f", composePath, "up", "-d")
	cmd.Stdout = nil
	cmd.Stderr = nil
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("ошибка запуска docker compose up: %w", err)
	}
	return nil
}

func (d *DockerComposeManager) Down(ctx context.Context) error {
	composePath := "docker-compose.yaml"
	cmd := exec.CommandContext(ctx, "docker", "compose", "-f", composePath, "down")
	cmd.Stdout = nil
	cmd.Stderr = nil
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("ошибка остановки docker compose down: %w", err)
	}
	return nil
}

func (d *DockerComposeManager) Reload(ctx context.Context) error {
	// Перегенерируем compose и перезапустим с --force-recreate
	composePath := "docker-compose.yaml"
	err := composegen.GenerateFromConfig(d.ConfigPath, composePath)
	if err != nil {
		return fmt.Errorf("ошибка генерации compose: %w", err)
	}
	cmd := exec.CommandContext(ctx, "docker", "compose", "-f", composePath, "up", "-d", "--force-recreate")
	cmd.Stdout = nil
	cmd.Stderr = nil
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("ошибка reload docker compose: %w", err)
	}
	return nil
}

func (d *DockerComposeManager) Status(ctx context.Context) (string, error) {
	composePath := "docker-compose.yaml"
	cmd := exec.CommandContext(ctx, "docker", "compose", "-f", composePath, "ps")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("ошибка получения статуса: %w", err)
	}
	return string(output), nil
}
