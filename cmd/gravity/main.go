package main

import (
	"context"
	"fmt"
	"log"
	"net"

	badger "github.com/dgraph-io/badger/v4"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"

	"github.com/vilahack/gravity/internal/bucket"
	"github.com/vilahack/gravity/internal/config"
	"github.com/vilahack/gravity/internal/file"
	pb "github.com/vilahack/gravity/internal/grpc"
	"github.com/vilahack/gravity/internal/server"
)

func main() {
	ctx := context.Background()

	cfg, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	db, err := badger.Open(badger.DefaultOptions(cfg.Database))
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	loggingShutdown, err := server.SetupLogging(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer loggingShutdown(ctx)

	bucketRepository := bucket.NewRepository(db)
	fileRepository := file.NewRepository(db)

	bucketService := bucket.NewService(bucketRepository)
	fileService := file.NewService(fileRepository)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.Port))
	if err != nil {
		log.Fatal(err)
	}

	var opts []grpc.ServerOption
	opts = append(opts, grpc.StatsHandler(otelgrpc.NewServerHandler()))

	s := grpc.NewServer(opts...)

	pb.RegisterBucketServiceServer(s, server.NewGRPCServer(bucketService, fileService))

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
