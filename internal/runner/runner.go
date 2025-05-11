package runner

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"text/template"

	"gopkg.in/yaml.v2"
)

type RunnerConfig struct {
	Runner      string `yaml:"runner"`
	Command     string `yaml:"command"`
	DownCommand string `yaml:"down_command"`
	Files       []struct {
		Template string `yaml:"template"`
		Output   string `yaml:"output"`
	} `yaml:"files"`
	Env []string `yaml:"env"`
}

// LoadRunnerConfig читает runner.yaml
func LoadRunnerConfig(path string) (*RunnerConfig, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var cfg RunnerConfig
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}

// RenderTemplates рендерит все шаблоны runner.yaml с подстановкой env
func RenderTemplates(cfg *RunnerConfig, env map[string]string, baseDir string) error {
	for _, f := range cfg.Files {
		tmplPath := filepath.Join(baseDir, f.Template)
		outPath := filepath.Join(baseDir, f.Output)
		tmplBytes, err := os.ReadFile(tmplPath)
		if err != nil {
			return fmt.Errorf("ошибка чтения шаблона %s: %w", tmplPath, err)
		}
		tmpl, err := template.New(f.Template).Parse(string(tmplBytes))
		if err != nil {
			return fmt.Errorf("ошибка парсинга шаблона %s: %w", tmplPath, err)
		}
		outFile, err := os.Create(outPath)
		if err != nil {
			return fmt.Errorf("ошибка создания файла %s: %w", outPath, err)
		}
		defer outFile.Close()
		if err := tmpl.Execute(outFile, env); err != nil {
			return fmt.Errorf("ошибка рендера шаблона %s: %w", tmplPath, err)
		}
	}
	return nil
}

// RunCommand запускает команду из runner.yaml с нужным env
func RunCommand(cfg *RunnerConfig, env map[string]string, baseDir string) error {
	cmd := exec.Command("bash", "-c", cfg.Command)
	cmd.Dir = baseDir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	// Подставляем env
	for k, v := range env {
		cmd.Env = append(cmd.Env, fmt.Sprintf("%s=%s", k, v))
	}
	cmd.Env = append(cmd.Env, os.Environ()...)
	return cmd.Run()
}

// RunDownCommand запускает down_command из runner.yaml с нужным env
func RunDownCommand(cfg *RunnerConfig, env map[string]string, baseDir string) error {
	if cfg.DownCommand == "" {
		return fmt.Errorf("down_command не задан в runner.yaml")
	}
	cmd := exec.Command("bash", "-c", cfg.DownCommand)
	cmd.Dir = baseDir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	for k, v := range env {
		cmd.Env = append(cmd.Env, fmt.Sprintf("%s=%s", k, v))
	}
	cmd.Env = append(cmd.Env, os.Environ()...)
	return cmd.Run()
}
