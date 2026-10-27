package server

import (
	"context"

	"github.com/vilahack/gravity/api/proto"
)

func (s *GRPCServer) UpsertBucket(
	_ context.Context,
	request *gravitypb.UpsertBucketRequest,
) (*gravitypb.UpsertBucketResponse, error) {
	return nil, nil // TODO
}
