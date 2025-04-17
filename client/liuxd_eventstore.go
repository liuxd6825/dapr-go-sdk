package client

import (
	"context"
	pb "github.com/dapr/dapr/pkg/proto/runtime/v1"
	_ "google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
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
	return c.protoClient.CommitDomainEvents(c.withAuthToken(ctx), request)
}

func (c *GRPCClient) RollbackDomainEvents(ctx context.Context, request *pb.RollbackDomainEventsRequest) (*pb.RollbackDomainEventsResponse, error) {
	return c.protoClient.RollbackDomainEvents(c.withAuthToken(ctx), request)
}

func (c *GRPCClient) GetDomainEventRelations(ctx context.Context, request *pb.GetDomainEventRelationsRequest) (*pb.GetDomainEventRelationsResponse, error) {
	return c.protoClient.GetDomainEventRelations(c.withAuthToken(ctx), request)
}

func (c *GRPCClient) GetDomainEvents(ctx context.Context, request *pb.GetDomainEventsRequest) (*pb.GetDomainEventsResponse, error) {
	return c.protoClient.GetDomainEvents(c.withAuthToken(ctx), request)
}

func (c *GRPCClient) withAuthToken(ctx context.Context) context.Context {
	if c.authToken == nil {
		return ctx
	}
	return metadata.NewOutgoingContext(ctx, metadata.Pairs(apiTokenKey, c.authToken.authToken))
}
