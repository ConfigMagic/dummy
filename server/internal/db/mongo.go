package db

import (
	"config_saver/internal/config"
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

type MongoDB struct {
	Client           *mongo.Client
	ConfigCollection *mongo.Collection
}

func NewMongoDB(cfg *config.AppConfig, logger *zap.Logger) (*MongoDB, error) {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(cfg.MongoDB.URI))
	if err != nil {
		logger.Error("Failed to connect to MongoDB", zap.Error(err))
		return nil, err
	}
	return &MongoDB{
		Client:           client,
		ConfigCollection: client.Database(cfg.MongoDB.DB).Collection(cfg.MongoDB.Collection),
	}, nil
}
