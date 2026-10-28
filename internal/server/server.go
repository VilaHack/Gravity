// Package server contains all the grpc connection handlers and loggers
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

	"github.com/vilahack/gravity/api/proto"
	"github.com/vilahack/gravity/internal/storage"
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
	gravitypb.UnimplementedBucketServiceServer
	StorageService *storage.Service
}

func NewGRPCServer(ss *storage.Service) *GRPCServer {
	return &GRPCServer{
		StorageService: ss,
	}
}
