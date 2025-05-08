package main

import (
	"fmt"
	"os"
	"io/ioutil"

	"github.com/spf13/cobra"
)

var pushCmd = &cobra.Command{
	Use:   "push [config-file]",
	Short: "Publish configuration to the server",
	Long:  "Read the configuration file and send it to the server (пока только выводит содержимое)",
	Args:  cobra.ExactArgs(1), // Требуем ровно один аргумент: путь к файлу
	Run: func(cmd *cobra.Command, args []string) {
		configPath := args[0]
		// Читаем файл конфигурации
		data, err := ioutil.ReadFile(configPath)
		if err != nil {
			fmt.Printf("Ошибка чтения файла %s: %v\n", configPath, err)
			os.Exit(1)
		}
		// Выводим содержимое файла (заглушка)
		fmt.Println("Содержимое конфигурационного файла:")
		fmt.Println(string(data))
		// TODO: Реализовать отправку на сервер
	},
}

func init() {
	rootCmd.AddCommand(pushCmd)
}
