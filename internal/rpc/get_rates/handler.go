package get_rates

import (
	"context"
	"github.com/pkg/errors"
	pb "testTask/grpc/rates"
)

type RateServiceServer struct {
	pb.UnimplementedRateServiceServer
	Usecase usecase
}

func New(usecase usecase) *RateServiceServer {
	return &RateServiceServer{
		Usecase: usecase,
	}
}

// GetRates handler grpc
func (s *RateServiceServer) GetRates(ctx context.Context, req *pb.RateRequest) (*pb.RateResponse, error) {
	if req.MarketID == "" {
		// todo log
		return nil, errors.New("empty marketID")
	}

	result, err := s.Usecase.Handle(ctx, req.MarketID)
	if err != nil {
		// todo log
		return nil, errors.Wrap(err, "get rates")
	}

	response := &pb.RateResponse{
		Timestamp: result.Timestamp,
		Ask:       result.AskPrice,
		Bid:       result.BidPrice,
	}

	return response, nil
}
