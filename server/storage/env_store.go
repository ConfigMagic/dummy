package storage

import (
	"fmt"
	"log"
	"os"

	"github.com/ConfigMagic/dummy/server/pb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var client *mongo.Client
var envCollection *mongo.Collection

// Initialize MongoDB connection
func InitMongo() error {
	uri := os.Getenv("MONGO_URI")
	if uri == "" {
		uri = "mongodb://localhost:27017" // Use local MongoDB by default
	}

	var err error
	client, err = mongo.Connect(nil, options.Client().ApplyURI(uri))
	if err != nil {
		return fmt.Errorf("failed to connect to MongoDB: %v", err)
	}

	// Check the connection
	if err := client.Ping(nil, readpref.Primary()); err != nil {
		return fmt.Errorf("failed to ping MongoDB: %v", err)
	}

	envCollection = client.Database("dummy").Collection("envs")
	return nil
}

// Close the MongoDB connection
func CloseMongo() {
	if err := client.Disconnect(nil); err != nil {
		log.Printf("failed to disconnect from MongoDB: %v", err)
	}
}

// Get environment configuration by name
func GetEnvConfig(name string) (*pb.EnvConfig, error) {
	var result pb.EnvConfig
	err := envCollection.FindOne(nil, bson.M{"name": name}).Decode(&result)
	if err != nil {
		return nil, fmt.Errorf("failed to find environment config: %v", err)
	}
	return &result, nil
}

// Get all environment configurations
func ListEnvConfigs() ([]*pb.EnvConfig, error) {
	cursor, err := envCollection.Find(nil, bson.D{})
	if err != nil {
		return nil, fmt.Errorf("failed to list environment configs: %v", err)
	}
	defer cursor.Close(nil)

	var envConfigs []*pb.EnvConfig
	for cursor.Next(nil) {
		var envConfig pb.EnvConfig
		if err := cursor.Decode(&envConfig); err != nil {
			return nil, fmt.Errorf("failed to decode environment config: %v", err)
		}
		envConfigs = append(envConfigs, &envConfig)
	}
	if err := cursor.Err(); err != nil {
		return nil, fmt.Errorf("cursor error: %v", err)
	}

	return envConfigs, nil
}

// Save environment configuration
func SaveEnvConfig(config *pb.EnvConfig) (*pb.EnvConfig, error) {
	_, err := envCollection.UpdateOne(
		nil,
		bson.M{"name": config.Name},
		bson.M{"$set": config},
		options.Update().SetUpsert(true),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to save environment config: %v", err)
	}
	return config, nil
}

// Delete environment configuration by name
func DeleteEnvConfig(name string) error {
	_, err := envCollection.DeleteOne(nil, bson.M{"name": name})
	if err != nil {
		return fmt.Errorf("failed to delete environment config: %v", err)
	}
	return nil
}
