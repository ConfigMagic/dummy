package main

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var pushCmd = &cobra.Command{
	Use:   "push [config-file]",
	Short: "Опубликовать конфигурацию на сервер (адрес через DUMMY_SERVER_URL)",
	Long: `Загружает или обновляет конфиг окружения на централизованном сервере для всех разработчиков.

Позволяет быстро распространять актуальные настройки и автоматизировать delivery конфигов. Сервер задаётся через переменную окружения DUMMY_SERVER_URL.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		configPath := args[0]

		// Имя конфига (без расширения)
		name := filepath.Base(configPath)
		if ext := filepath.Ext(name); ext == ".yaml" || ext == ".yml" {
			name = name[:len(name)-len(ext)]
		}

		// Папка с файлами (рядом с ямлом, с тем же именем)
		configDir := configPath[:len(configPath)-len(filepath.Ext(configPath))]
		var buf bytes.Buffer
		zipWriter := zip.NewWriter(&buf)

		// Добавляем сам YAML
		yamlFile, err := os.Open(configPath)
		if err != nil {
			exitWithError(fmt.Errorf("ошибка чтения yaml: %v", err))
		}
		defer yamlFile.Close()
		w, err := zipWriter.Create("config.yaml")
		if err != nil {
			exitWithError(fmt.Errorf("ошибка zip: %v", err))
		}
		io.Copy(w, yamlFile)

		// Если есть папка с файлами — добавляем их в корень архива (без files/)
		if fi, err := os.Stat(configDir); err == nil && fi.IsDir() {
			filepath.Walk(configDir, func(path string, info os.FileInfo, err error) error {
				if err != nil || info.IsDir() {
					return nil
				}
				rel, _ := filepath.Rel(configDir, path)
				// Не добавлять сам config.yaml второй раз
				if rel == "" || rel == filepath.Base(configPath) {
					return nil
				}
				f, err := os.Open(path)
				if err != nil {
					return nil
				}
				defer f.Close()
				w, err := zipWriter.Create(rel)
				if err != nil {
					return nil
				}
				io.Copy(w, f)
				return nil
			})
		}
		zipWriter.Close()

		serverURL := os.Getenv("DUMMY_SERVER_URL")
		if serverURL == "" {
			serverURL = "http://localhost:8080"
		}
		url := serverURL + "/config-multipart"

		body := &bytes.Buffer{}
		writer := multipart.NewWriter(body)
		writer.WriteField("name", name)
		fw, err := writer.CreateFormFile("archive", name+".zip")
		if err != nil {
			exitWithError(fmt.Errorf("ошибка multipart: %v", err))
		}
		fw.Write(buf.Bytes())
		writer.Close()

		resp, err := http.Post(url, writer.FormDataContentType(), body)
		if err != nil {
			exitWithError(fmt.Errorf("ошибка отправки POST: %v", err))
		}
		defer resp.Body.Close()

		if resp.StatusCode != 200 {
			b, _ := io.ReadAll(resp.Body)
			exitWithError(fmt.Errorf("сервер вернул ошибку: %s", string(b)))
		}

		fmt.Printf("✅ Конфиг '%s' (yaml+files) успешно загружен на сервер\n", name)
	},
}

func init() {
	rootCmd.AddCommand(pushCmd)
}
