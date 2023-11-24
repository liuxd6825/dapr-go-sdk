package client

import (
	"context"
	pb "github.com/liuxd6825/dapr/pkg/proto/runtime/v1"
	_ "google.golang.org/grpc"
)

// LoadDomainEvents
// @Description:  DDD event storage get aggregate root by id
// @receiver c
// @param ctx
// @param req
// @return *pb.LoadEventResponse
// @return error
func (c *GRPCClient) LoadDomainEvents(ctx context.Context, req *pb.LoadDomainEventRequest) (*pb.LoadDomainEventResponse, error) {
	return c.protoClient.LoadDomainEvent(c.withAuthToken(ctx), req)
}

// SaveDomainEventSnapshot
// @Description: liuxd: DDD event storage apply event
// @receiver c
// @param ctx
// @param in
// @return *pb.SaveSnapshotResponse
// @return error
func (c *GRPCClient) SaveDomainEventSnapshot(ctx context.Context, in *pb.SaveDomainEventSnapshotRequest) (*pb.SaveDomainEventSnapshotResponse, error) {
	return c.protoClient.SaveDomainEventSnapshot(c.withAuthToken(ctx), in)
}

// ApplyDomainEvent
// @Description: liuxd: DDD event storage apply event
// @receiver c
// @param ctx
// @param in
// @return *pb.ApplyEventsResponse
// @return error
func (c *GRPCClient) ApplyDomainEvent(ctx context.Context, in *pb.ApplyDomainEventRequest) (*pb.ApplyDomainEventResponse, error) {
	return c.protoClient.ApplyDomainEvent(c.withAuthToken(ctx), in)
}

func (c *GRPCClient) CommitDomainEvents(ctx context.Context, request *pb.CommitDomainEventsRequest) (*pb.CommitDomainEventsResponse, error) {
	return c.protoClient.CommitDomainEvents(ctx, request)
}

func (c *GRPCClient) RollbackDomainEvents(ctx context.Context, request *pb.RollbackDomainEventsRequest) (*pb.RollbackDomainEventsResponse, error) {
	return c.protoClient.RollbackDomainEvents(ctx, request)
}

func (c *GRPCClient) GetDomainEventRelations(ctx context.Context, request *pb.GetDomainEventRelationsRequest) (*pb.GetDomainEventRelationsResponse, error) {
	return c.protoClient.GetDomainEventRelations(ctx, request)
}

func (c *GRPCClient) GetDomainEvents(ctx context.Context, request *pb.GetDomainEventsRequest) (*pb.GetDomainEventsResponse, error) {
	return c.protoClient.GetDomainEvents(ctx, request)
}
