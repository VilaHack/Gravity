package server

import (
	"context"

	"github.com/vilahack/gravity/api/proto"
)

func (s *GRPCServer) ListFiles(
	_ context.Context,
	request *gravitypb.ListFilesRequest,
) (*gravitypb.ListFilesResponse, error) {
	return nil, nil // TODO
}

func (s *GRPCServer) DownloadFile(
	request *gravitypb.DownloadFileRequest,
	stream gravitypb.BucketService_DownloadFileServer,
) error {
	return nil // TODO
}

func (s *GRPCServer) UploadFile(
	stream gravitypb.BucketService_UploadFileServer,
) error {
	return nil // TODO
}
