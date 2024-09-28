package healthcheck

import (
	"context"
	pb "testTask/grpc/healthcheck"
)

type HealthCheckServer struct {
	pb.UnimplementedHealthcheckServiceServer
}

func New() *HealthCheckServer {
	return &HealthCheckServer{}
}

// GetRates handler grpc
func (h *HealthCheckServer) Healthcheck(ctx context.Context, req *pb.HealthcheckRequest) (*pb.HealthcheckResponse, error) {
	return &pb.HealthcheckResponse{
		Message: "health",
	}, nil
}
