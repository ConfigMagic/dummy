package main

import (
	"archive/zip"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

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
		url := fmt.Sprintf("%s/config-multipart?name=%s", serverURL, configName)

		resp, err := http.Get(url)
		if err != nil {
			exitWithError(fmt.Errorf("ошибка запроса к серверу: %v", err))
		}
		defer resp.Body.Close()

		if resp.StatusCode != 200 {
			body, _ := io.ReadAll(resp.Body)
			exitWithError(fmt.Errorf("сервер вернул ошибку: %s", string(body)))
		}

		// Сохраняем архив во временный файл
		tmpZip := filepath.Join(os.TempDir(), configName+".zip")
		f, err := os.Create(tmpZip)
		if err != nil {
			exitWithError(fmt.Errorf("ошибка создания временного файла: %v", err))
		}
		io.Copy(f, resp.Body)
		f.Close()

		// Распаковываем архив в ~/.dummy/<configName>/
		homeDir, err := os.UserHomeDir()
		if err != nil {
			exitWithError(fmt.Errorf("не удалось определить домашний каталог: %v", err))
		}
		destDir := filepath.Join(homeDir, ".dummy", configName)
		os.MkdirAll(destDir, 0755)
		zipReader, err := zip.OpenReader(tmpZip)
		if err != nil {
			exitWithError(fmt.Errorf("ошибка открытия архива: %v", err))
		}
		defer zipReader.Close()
		for _, f := range zipReader.File {
			fpath := filepath.Join(destDir, f.Name)
			if f.FileInfo().IsDir() {
				os.MkdirAll(fpath, f.Mode())
				continue
			}
			os.MkdirAll(filepath.Dir(fpath), 0755)
			outFile, err := os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
			if err != nil {
				continue
			}
			inFile, err := f.Open()
			if err != nil {
				outFile.Close()
				continue
			}
			io.Copy(outFile, inFile)
			outFile.Close()
			inFile.Close()
		}
		os.Remove(tmpZip)

		fmt.Printf("✅ Конфиг '%s' (yaml+files) синхронизирован и распакован в %s\n", configName, destDir)
	},
}

func init() {
	rootCmd.AddCommand(syncCmd)
}
