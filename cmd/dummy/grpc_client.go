package main

import (
	"context"
	"fmt"

	pb "github.com/ConfigMagic/dummy/server/pb"
	"google.golang.org/grpc"
)

type GRPCClient struct {
	conn   *grpc.ClientConn
	client pb.EnvServiceClient
}

func NewGRPCClient(serverAddr string) (*GRPCClient, error) {
	conn, err := grpc.Dial(serverAddr, grpc.WithInsecure())
	if err != nil {
		return nil, fmt.Errorf("не удалось подключиться к серверу: %v", err)
	}
	return &GRPCClient{
		conn:   conn,
		client: pb.NewEnvServiceClient(conn),
	}, nil
}

// Получить конфиг по имени
func (c *GRPCClient) GetConfig(ctx context.Context, name string) (*pb.EnvConfig, error) {
	req := &pb.EnvRequest{Name: name}
	return c.client.GetEnvConfig(ctx, req)
}

// Список всех конфигов
func (c *GRPCClient) ListConfigs(ctx context.Context) (*pb.EnvList, error) {
	req := &pb.Empty{}
	return c.client.ListEnvConfigs(ctx, req)
}

// Загрузить новый конфиг
func (c *GRPCClient) ApplyConfig(ctx context.Context, config *pb.EnvConfig) (*pb.EnvConfig, error) {
	return c.client.ApplyEnvConfig(ctx, config)
}

// Удалить конфиг
func (c *GRPCClient) DeleteConfig(ctx context.Context, name string) (*pb.Empty, error) {
	req := &pb.EnvRequest{Name: name}
	return c.client.DeleteEnvConfig(ctx, req)
}

func (c *GRPCClient) Close() {
	c.conn.Close()
}
