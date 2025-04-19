package main

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ConfigMagic/dummy/server/handlers"
	"github.com/ConfigMagic/dummy/server/pb"
	"github.com/ConfigMagic/dummy/server/storage"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// Server holds the gRPC server, logger, and MongoDB connection
type Server struct {
	grpcServer *grpc.Server
	logger     *zap.SugaredLogger
}

// NewServer initializes a new Server instance
func NewServer() (*Server, error) {
	// Initialize logger
	logger, err := initLogger()
	if err != nil {
		return nil, fmt.Errorf("failed to initialize logger: %v", err)
	}

	// Initialize MongoDB with retries
	if err := initMongoWithRetry(logger); err != nil {
		logger.Errorw("Mongo connection failed", "error", err)
		return nil, err
	}

	// Create gRPC server
	grpcServer := grpc.NewServer()
	pb.RegisterEnvServiceServer(grpcServer, &handlers.EnvServiceServer{})
	reflection.Register(grpcServer)

	return &Server{
		grpcServer: grpcServer,
		logger:     logger,
	}, nil
}

// Start runs the gRPC server
func (s *Server) Start(port string) error {
	if port == "" {
		port = "50051"
	}

	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		s.logger.Fatalf("failed to listen: %v", err)
	}

	s.logger.Infof("gRPC server started on port %s", port)

	// Handle shutdown signals
	go s.handleShutdown()

	// Start serving
	return s.grpcServer.Serve(listener)
}

// Stop gracefully stops the gRPC server and closes resources
func (s *Server) Stop() {
	s.logger.Info("Stopping server...")
	s.grpcServer.GracefulStop()
	storage.CloseMongo()
	s.logger.Sync()
}

// handleShutdown listens for OS signals and stops the server
func (s *Server) handleShutdown() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	<-c
	s.logger.Info("Received shutdown signal. Stopping...")
	s.Stop()
}

// initLogger initializes the zap logger
func initLogger() (*zap.SugaredLogger, error) {
	logDir := "/var/log/dummy-admin"
	logFile := fmt.Sprintf("%s/server.log", logDir)

	if err := os.MkdirAll(logDir, 0755); err != nil {
		return nil, fmt.Errorf("failed to create log directory: %v", err)
	}

	cfg := zap.Config{
		Encoding:         "json",
		Level:            zap.NewAtomicLevelAt(zap.InfoLevel),
		OutputPaths:      []string{"stdout", logFile},
		ErrorOutputPaths: []string{"stderr", logFile},
		EncoderConfig:    zap.NewProductionEncoderConfig(),
	}

	logger, err := cfg.Build()
	if err != nil {
		return nil, fmt.Errorf("failed to initialize logger: %v", err)
	}
	return logger.Sugar(), nil
}

// initMongoWithRetry initializes MongoDB with retry logic
func initMongoWithRetry(logger *zap.SugaredLogger) error {
	const maxRetries = 10
	const delay = 2 * time.Second

	for i := 1; i <= maxRetries; i++ {
		err := storage.InitMongo()
		if err == nil {
			logger.Infow("Connected to MongoDB", "attempt", i)
			return nil
		}
		logger.Warnw("MongoDB not ready, retrying...", "attempt", i, "error", err)
		time.Sleep(delay)
	}
	return fmt.Errorf("failed to connect to MongoDB after %d attempts", maxRetries)
}

func main() {
	// Initialize the server
	server, err := NewServer()
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to initialize server: %v\n", err)
		os.Exit(1)
	}

	// Start the server
	port := os.Getenv("PORT")
	if err := server.Start(port); err != nil {
		server.logger.Fatalf("failed to serve: %v", err)
	}
}
