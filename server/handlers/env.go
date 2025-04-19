package handlers

import (
	"context"
	"fmt"

	"github.com/ConfigMagic/dummy/server/pb"
	"github.com/ConfigMagic/dummy/server/storage"
)

type EnvServiceServer struct {
	pb.UnimplementedEnvServiceServer
}

// Method to get the environment configuration by name
func (s *EnvServiceServer) GetEnvConfig(ctx context.Context, req *pb.EnvRequest) (*pb.EnvConfig, error) {
	envConfig, err := storage.GetEnvConfig(req.GetName())
	if err != nil {
		return nil, fmt.Errorf("failed to get environment config: %v", err)
	}
	return envConfig, nil
}

// Method to get all environment configurations
func (s *EnvServiceServer) ListEnvConfigs(ctx context.Context, req *pb.Empty) (*pb.EnvList, error) {
	envConfigs, err := storage.ListEnvConfigs()
	if err != nil {
		return nil, fmt.Errorf("failed to list environment configs: %v", err)
	}
	return &pb.EnvList{Configs: envConfigs}, nil
}

// Method to create or update an environment configuration
func (s *EnvServiceServer) ApplyEnvConfig(ctx context.Context, req *pb.EnvConfig) (*pb.EnvConfig, error) {
	// Logic to save to the database
	envConfig, err := storage.SaveEnvConfig(req)
	if err != nil {
		return nil, fmt.Errorf("failed to apply environment config: %v", err)
	}
	return envConfig, nil
}

// Method to delete an environment configuration
func (s *EnvServiceServer) DeleteEnvConfig(ctx context.Context, req *pb.EnvRequest) (*pb.Empty, error) {
	if err := storage.DeleteEnvConfig(req.GetName()); err != nil {
		return nil, fmt.Errorf("failed to delete environment config: %v", err)
	}
	return &pb.Empty{}, nil
}
