package client

import (
	"context"
	pb "github.com/liuxd6825/dapr/pkg/proto/runtime/v1"
)

func (c *GRPCClient) Command(ctx context.Context, in *pb.CommandRequest) (*pb.CommandResponse, error) {
	return c.protoClient.Command(c.withAuthToken(ctx), in)
}
