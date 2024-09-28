package get_rates

import (
	"context"
	pb "testTask/grpc/rates"
)

type rateServiceServer struct {
	pb.UnimplementedRateServiceServer
}

// GetRates handler grpc
func (s *rateServiceServer) GetRates(ctx context.Context, req *pb.RateRequest) (*pb.RateResponse, error) {
	return nil, nil
}
