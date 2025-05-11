package main

import (
	"context"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"config_saver/internal/config"
	"config_saver/internal/db"
	"config_saver/internal/handler"
	"config_saver/internal/logger"
)

func main() {
	// Логгер
	log := logger.New()
	defer log.Sync()

	// Конфиг из ENV
	cfg := config.LoadFromEnv()

	// MongoDB
	mongoDB, err := db.NewMongoDB(cfg, log)
	if err != nil {
		log.Fatal("Failed to init MongoDB", zap.Error(err))
	}
	defer mongoDB.Client.Disconnect(context.Background())
	db.RunMigrations(mongoDB.Client.Database(cfg.MongoDB.DB), log)

	// Gin
	r := gin.Default()
	configHandler := &handler.ConfigHandler{DB: mongoDB, Logger: log}

	// Роуты
	r.GET("/config/:name", configHandler.GetConfig)
	r.POST("/config/:name", configHandler.SaveConfig)
	r.DELETE("/config/:name", configHandler.DeleteConfig)

	// Запуск сервера
	log.Info("Server started", zap.String("port", cfg.Server.Port))
	r.Run(":" + cfg.Server.Port)
}
