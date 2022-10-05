package grpc

import (
	"context"
	"log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/rubenvanstaden/grpc-health/pb"
)

type HealthService struct {
	statusMap map[string]pb.HealthCheckResponse_ServingStatus
	pb.UnimplementedHealthServer
}

func NewHealthService() *HealthService {
	return &HealthService{
		statusMap: map[string]pb.HealthCheckResponse_ServingStatus{
			"": pb.HealthCheckResponse_SERVING,
		},
	}
}

func (self *HealthService) Check(ctx context.Context, request *pb.HealthCheckRequest) (*pb.HealthCheckResponse, error) {
	s, ok := self.statusMap[request.Service]
	if !ok {
		return nil, status.Error(codes.NotFound, "unknown service")
	}
	return &pb.HealthCheckResponse{Status: s}, nil
}

func (s *HealthService) Watch(in *pb.HealthCheckRequest, request pb.Health_WatchServer) error {
	op := "health.Watch"
	log.Fatalf("Not Implemented: %s", op)
	return nil
}
