package healthcheck

import (
	"context"
	pb "testTask/grpc/healthcheck"
)

type Server struct {
	pb.UnimplementedHealthcheckServiceServer
}

func New() *Server {
	return &Server{}
}

// GetRates handler grpc
func (h *Server) Healthcheck(context.Context, *pb.HealthcheckRequest) (*pb.HealthcheckResponse, error) {
	return &pb.HealthcheckResponse{
		Message: "health",
	}, nil
}
