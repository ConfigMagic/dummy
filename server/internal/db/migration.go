package db

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

func RunMigrations(db *mongo.Database, log *zap.Logger) {
	log.Debug("Running migrations")
	_, err := db.Collection("configs").Indexes().CreateOne(
		context.Background(),
		mongo.IndexModel{
			Keys:    bson.D{{Key: "name", Value: 1}},
			Options: nil,
		},
	)
	if err != nil {
		log.Error("Migration failed", zap.Error(err))
	} else {
		log.Info("Migrations applied")
	}
}
