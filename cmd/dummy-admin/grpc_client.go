package main

import (
	"context"
	"fmt"

	pb "github.com/ConfigMagic/dummy/server/pb"
	"google.golang.org/grpc"
)

type GRPCAdminClient struct {
	conn   *grpc.ClientConn
	client pb.EnvServiceClient
}

func NewGRPCAdminClient(serverAddr string) (*GRPCAdminClient, error) {
	conn, err := grpc.Dial(serverAddr, grpc.WithInsecure())
	if err != nil {
		return nil, fmt.Errorf("не удалось подключиться к серверу: %v", err)
	}
	return &GRPCAdminClient{
		conn:   conn,
		client: pb.NewEnvServiceClient(conn),
	}, nil
}

// Загрузить конфиг на сервер
func (c *GRPCAdminClient) ApplyConfig(ctx context.Context, config *pb.EnvConfig) (*pb.EnvConfig, error) {
	return c.client.ApplyEnvConfig(ctx, config)
}

// Управление пользователями
func (c *GRPCAdminClient) ManageUser(ctx context.Context, action string, username string) error {
	// Логика взаимодействия с сервером
	fmt.Printf("Выполняется действие: %s для пользователя: %s\n", action, username)
	return nil
}

func (c *GRPCAdminClient) Close() {
	c.conn.Close()
}
