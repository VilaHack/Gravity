package server

import (
	"context"

	"github.com/vilahack/gravity/api/proto"
)

func (s *GRPCServer) ListObjects(
	_ context.Context,
	request *gravitypb.ListObjectsRequest,
) (*gravitypb.ListObjectsResponse, error) {
	return nil, nil // TODO
}

func (s *GRPCServer) DownloadObject(
	request *gravitypb.DownloadObjectRequest,
	stream gravitypb.BucketService_DownloadObjectServer,
) error {
	return nil // TODO
}

func (s *GRPCServer) UploadObject(
	stream gravitypb.BucketService_UploadObjectServer,
) error {
	return nil // TODO
}
