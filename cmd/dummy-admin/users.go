package main

import "github.com/spf13/cobra"

var usersCmd = &cobra.Command{
	Use:   "users",
	Short: "Manage users",
	Long:  "Manage users: add, list, remove.",
}

// Подкоманда для добавления пользователя
var usersAddCmd = &cobra.Command{
	Use:   "add [username]",
	Short: "Add a new user",
	Long:  "Add a new user (заглушка)",
	Args:  cobra.ExactArgs(1), // Требуем ровно один аргумент: имя пользователя
	Run: func(cmd *cobra.Command, args []string) {
		username := args[0]
		// Здесь будет логика добавления пользователя
		cmd.Printf("Пользователь '%s' добавлен (заглушка)\n", username)
	},
}

// Подкоманда для вывода списка пользователей
var usersListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all users",
	Long:  "List all users (заглушка)",
	Run: func(cmd *cobra.Command, args []string) {
		// Здесь будет логика вывода списка пользователей
		cmd.Println("Список пользователей: user1, user2 (заглушка)")
	},
}

// Подкоманда для удаления пользователя
var usersRemoveCmd = &cobra.Command{
	Use:   "remove [username]",
	Short: "Remove a user",
	Long:  "Remove a user (заглушка)",
	Args:  cobra.ExactArgs(1), // Требуем ровно один аргумент: имя пользователя
	Run: func(cmd *cobra.Command, args []string) {
		username := args[0]
		// Здесь будет логика удаления пользователя
		cmd.Printf("Пользователь '%s' удалён (заглушка)\n", username)
	},
}

func init() {
	// Добавляем подкоманды к usersCmd
	usersCmd.AddCommand(usersAddCmd)
	usersCmd.AddCommand(usersListCmd)
	usersCmd.AddCommand(usersRemoveCmd)
	// Добавляем usersCmd к rootCmd
	rootCmd.AddCommand(usersCmd)
}
