package main

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	badger "github.com/dgraph-io/badger/v4"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"

	"github.com/vilahack/gravity/api/proto"
	"github.com/vilahack/gravity/internal/bucket"
	"github.com/vilahack/gravity/internal/config"
	"github.com/vilahack/gravity/internal/file"
	"github.com/vilahack/gravity/internal/server"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	cfg, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	db, err := badger.Open(badger.DefaultOptions(cfg.Database))
	if err != nil {
		slog.Error("failed to open database", "error", err)
		os.Exit(1)
	}

	defer db.Close()

	loggingShutdown, err := server.SetupLogging(ctx)
	if err != nil {
		slog.Error("failed to setup logging", "error", err)
		os.Exit(1)
	}

	defer func() {
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := loggingShutdown(shutdownCtx); err != nil {
			fmt.Fprintf(os.Stderr, "failed to shutdown logging: %v\n", err)
		}
	}()

	slog.Info("starting gravity microservice", "port", cfg.Port)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.Port))
	if err != nil {
		slog.Error("failed to start listener", "error", err)
		os.Exit(1)
	}

	var opts []grpc.ServerOption
	opts = append(opts, grpc.StatsHandler(otelgrpc.NewServerHandler()))
	s := grpc.NewServer(opts...)

	bucketRepository := bucket.NewRepository(db)
	fileRepository := file.NewRepository(db)

	bucketService := bucket.NewService(bucketRepository)
	fileService := file.NewService(fileRepository)

	gravitypb.RegisterBucketServiceServer(s, server.NewGRPCServer(bucketService, fileService))

	serverErr := make(chan error, 1)
	go func() {
		if err := s.Serve(lis); err != nil && !errors.Is(err, grpc.ErrServerStopped) {
			serverErr <- err
		}
	}()

	select {
	case err := <-serverErr:
		slog.Error("failed to start server", "error", err)

	case <-ctx.Done():
		slog.Info("shutting down server...")
		s.GracefulStop()
	}
}
