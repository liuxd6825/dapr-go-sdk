package client

import (
	"context"
	pb "github.com/liuxd6825/dapr/pkg/proto/runtime/v1"
)

func (c *GRPCClient) WriteAppEventLog(ctx context.Context, in *pb.WriteAppEventLogRequest) (*pb.WriteAppEventLogResponse, error) {
	return c.protoClient.WriteAppEventLog(c.withAuthToken(ctx), in)
}

func (c *GRPCClient) UpdateAppEventLog(ctx context.Context, in *pb.UpdateAppEventLogRequest) (*pb.UpdateAppEventLogResponse, error) {
	return c.protoClient.UpdateAppEventLog(c.withAuthToken(ctx), in)
}

func (c *GRPCClient) GetAppEventLogByCommandId(ctx context.Context, in *pb.GetAppEventLogByCommandIdRequest) (*pb.GetAppEventLogByCommandIdResponse, error) {
	return c.protoClient.GetAppEventLogByCommandId(c.withAuthToken(ctx), in)
}

func (c *GRPCClient) WriteAppLog(ctx context.Context, in *pb.WriteAppLogRequest) (*pb.WriteAppLogResponse, error) {
	return c.protoClient.WriteAppLog(c.withAuthToken(ctx), in)
}

func (c *GRPCClient) UpdateAppLog(ctx context.Context, in *pb.UpdateAppLogRequest) (*pb.UpdateAppLogResponse, error) {
	return c.protoClient.UpdateAppLog(c.withAuthToken(ctx), in)
}

func (c *GRPCClient) GetAppLogById(ctx context.Context, in *pb.GetAppLogByIdRequest) (*pb.GetAppLogByIdResponse, error) {
	return c.protoClient.GetAppLogById(c.withAuthToken(ctx), in)
}
