package getrates

import (
	"context"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	pb "testTask/grpc/rates"
)

type RateServiceServer struct {
	pb.UnimplementedRateServiceServer
	Usecase usecase
	logger  *zap.SugaredLogger
}

func New(usecase usecase, logger *zap.SugaredLogger) *RateServiceServer {
	return &RateServiceServer{
		Usecase: usecase,
		logger:  logger,
	}
}

// GetRates handler grpc
func (s *RateServiceServer) GetRates(ctx context.Context, req *pb.RateRequest) (*pb.RateResponse, error) {
	if req.MarketID == "" {
		errorMessage := "empty marketID"
		s.logger.Errorw(errorMessage)

		return nil, errors.New(errorMessage)
	}

	result, err := s.Usecase.Handle(ctx, req.MarketID)
	if err != nil {
		s.logger.Errorw("failed get rates",
			"error", err,
		)

		return nil, errors.Wrap(err, "get rates")
	}

	response := &pb.RateResponse{
		Timestamp: result.Timestamp,
		Ask:       result.AskPrice,
		Bid:       result.BidPrice,
	}

	return response, nil
}
