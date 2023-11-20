package client

import (
	"context"
	pb "github.com/liuxd6825/dapr/pkg/proto/runtime/v1"
	_ "google.golang.org/grpc"
)

// LoadEvents
// @Description:  DDD event storage get aggregate root by id
// @receiver c
// @param ctx
// @param req
// @return *pb.LoadEventResponse
// @return error
func (c *GRPCClient) LoadEvents(ctx context.Context, req *pb.LoadEventRequest) (*pb.LoadEventResponse, error) {
	return c.protoClient.LoadEvents(c.withAuthToken(ctx), req)
}

// SaveSnapshot
// @Description: liuxd: DDD event storage apply event
// @receiver c
// @param ctx
// @param in
// @return *pb.SaveSnapshotResponse
// @return error
func (c *GRPCClient) SaveSnapshot(ctx context.Context, in *pb.SaveSnapshotRequest) (*pb.SaveSnapshotResponse, error) {
	return c.protoClient.SaveSnapshot(c.withAuthToken(ctx), in)
}

// ApplyEvent
// @Description: liuxd: DDD event storage apply event
// @receiver c
// @param ctx
// @param in
// @return *pb.ApplyEventsResponse
// @return error
func (c *GRPCClient) ApplyEvent(ctx context.Context, in *pb.ApplyEventRequest) (*pb.ApplyEventResponse, error) {
	return c.protoClient.ApplyEvent(c.withAuthToken(ctx), in)
}

func (c *GRPCClient) Commit(ctx context.Context, request *pb.CommitRequest) (*pb.CommitResponse, error) {
	return c.protoClient.Commit(ctx, request)
}

func (c *GRPCClient) Rollback(ctx context.Context, request *pb.RollbackRequest) (*pb.RollbackResponse, error) {
	return c.protoClient.Rollback(ctx, request)
}

func (c *GRPCClient) GetRelations(ctx context.Context, request *pb.GetRelationsRequest) (*pb.GetRelationsResponse, error) {
	return c.protoClient.GetRelations(ctx, request)
}

func (c *GRPCClient) GetEvents(ctx context.Context, request *pb.GetEventsRequest) (*pb.GetEventsResponse, error) {
	return c.protoClient.GetEvents(ctx, request)
}
