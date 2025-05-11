package logger

import (
	"os"
	"path/filepath"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func New() *zap.Logger {
	// Создаем директорию для логов если её нет
	logDir := "/var/log/dummy-admin"
	if err := os.MkdirAll(logDir, 0755); err != nil {
		panic(err)
	}

	// Настраиваем вывод в файл
	file, err := os.OpenFile(
		filepath.Join(logDir, "server.log"),
		os.O_APPEND|os.O_CREATE|os.O_WRONLY,
		0644,
	)
	if err != nil {
		panic(err)
	}

	// Создаем конфигурацию для логгера
	config := zap.NewProductionConfig()
	config.EncoderConfig.TimeKey = "timestamp"
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	// Создаем core для вывода в файл
	fileCore := zapcore.NewCore(
		zapcore.NewJSONEncoder(config.EncoderConfig),
		zapcore.AddSync(file),
		zapcore.InfoLevel,
	)

	// Создаем core для вывода в консоль
	consoleCore := zapcore.NewCore(
		zapcore.NewConsoleEncoder(config.EncoderConfig),
		zapcore.AddSync(os.Stdout),
		zapcore.InfoLevel,
	)

	// Объединяем cores
	core := zapcore.NewTee(fileCore, consoleCore)

	// Создаем логгер
	logger := zap.New(core)
	return logger
}
