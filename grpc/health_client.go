package grpc

import (
	"context"
	"log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/rubenvanstaden/grpc-health/grpc/pb"

	"github.com/rubenvanstaden/grpc-health/core"
)

type HealthClient struct {
	client pb.HealthClient
	flags  core.Flags
}

func NewHealthClient(client *Client) *HealthClient {
	return &HealthClient{
		client: pb.NewHealthClient(client.Open()),
		flags:  client.flags,
	}
}

func (self *HealthClient) Check(ctx context.Context, service string) int {

	request := &pb.HealthCheckRequest{
		Service: service,
	}

	response, err := self.client.Check(ctx, request)
	if err != nil {
		if stat, ok := status.FromError(err); ok && stat.Code() == codes.Unimplemented {
			log.Printf("error: this server does not implement the grpc health protocol (grpc.health.v1.Health): %s", stat.Message())
		} else if stat, ok := status.FromError(err); ok && stat.Code() == codes.DeadlineExceeded {
			log.Printf("timeout: health rpc did not complete within %v", self.flags.RPCTimeout)
		} else {
			log.Printf("error: health rpc failed: %+v", err)
		}
		return core.StatusRPCFailure
	}

	if response.GetStatus() != pb.HealthCheckResponse_SERVING {
		log.Printf("service unhealthy (responded with %q)", response.GetStatus())
		return core.StatusUnhealthy
	}

	return core.StatusSuccess
}

func (self *HealthClient) Watch(ctx context.Context, service string) (string, error) {
	op := "health.Watch"
	log.Fatalf("Not Implemented: %s", op)
	return "", nil
}
