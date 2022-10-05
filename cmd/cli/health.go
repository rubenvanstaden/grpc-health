package main

import (
	"context"
	"log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/rubenvanstaden/grpc-health/pb"
)

type HealthClient struct {
	client pb.HealthClient
}

func NewHealthClient(client *Client) *HealthClient {
	return &HealthClient{
		client: pb.NewHealthClient(client.Open()),
	}
}

func (self *HealthClient) Check(ctx context.Context, service string) {

	request := &pb.HealthCheckRequest{
		Service: service,
	}

	response, err := self.client.Check(ctx, request)
	if err != nil {
		if stat, ok := status.FromError(err); ok && stat.Code() == codes.Unimplemented {
			log.Printf("error: this server does not implement the grpc health protocol (grpc.health.v1.Health): %s", stat.Message())
		} else if stat, ok := status.FromError(err); ok && stat.Code() == codes.DeadlineExceeded {
			log.Printf("timeout: health rpc did not complete within %v", flRPCTimeout)
		} else {
			log.Printf("error: health rpc failed: %+v", err)
		}
		statusCode = StatusRPCFailure
	}

	if response.GetStatus() != pb.HealthCheckResponse_SERVING {
		log.Printf("service unhealthy (responded with %q)", response.GetStatus())
		statusCode = StatusUnhealthy
	}
}

func (self *HealthClient) Watch(ctx context.Context, service string) (string, error) {
	op := "health.Watch"
	log.Fatalf("Not Implemented: %s", op)
	return "", nil
}
