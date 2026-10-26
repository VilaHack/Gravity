// Package server contains all the grpc connection handlers
package server

import (
	"context"
	"fmt"
	"log/slog"

	"go.opentelemetry.io/contrib/bridges/otelslog"
	"go.opentelemetry.io/otel/exporters/otlp/otlplog/otlploggrpc"
	sdklog "go.opentelemetry.io/otel/sdk/log"
	"go.opentelemetry.io/otel/sdk/resource"
	semconv "go.opentelemetry.io/otel/semconv/v1.37.0"

	"github.com/vilahack/gravity/internal/bucket"
	"github.com/vilahack/gravity/internal/file"
	pb "github.com/vilahack/gravity/internal/grpc"
)

func SetupLogging(ctx context.Context) (func(context.Context) error, error) {
	exporter, err := otlploggrpc.New(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to create log exporter: %w", err)
	}

	res, err := resource.Merge(
		resource.Default(),
		resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceName("gravity"),
		),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create resource: %w", err)
	}

	provider := sdklog.NewLoggerProvider(
		sdklog.WithResource(res),
		sdklog.WithProcessor(sdklog.NewBatchProcessor(exporter)),
	)

	slog.SetDefault(otelslog.NewLogger(
		"gravity",
		otelslog.WithLoggerProvider(provider),
	))

	return provider.Shutdown, nil
}

type GRPCServer struct {
	pb.UnimplementedBucketServiceServer
	BucketService *bucket.Service
	FileService   *file.Service
}

func NewGRPCServer(bucketService *bucket.Service, fileService *file.Service) *GRPCServer {
	return &GRPCServer{
		BucketService: bucketService,
		FileService:   fileService,
	}
}

func (s *GRPCServer) ListFiles(
	_ context.Context,
	request *pb.ListFilesRequest,
) (*pb.ListFilesResponse, error) {
	return nil, nil // TODO
}

func (s *GRPCServer) DownloadFile(
	request *pb.DownloadFileRequest,
	stream pb.BucketService_DownloadFileServer,
) error {
	return nil // TODO
}

func (s *GRPCServer) UploadFile(
	stream pb.BucketService_UploadFileServer,
) error {
	return nil // TODO
}

func (s *GRPCServer) UpsertBucket(
	_ context.Context,
	request *pb.UpsertBucketRequest,
) (*pb.UpsertBucketResponse, error) {
	return nil, nil // TODO
}
