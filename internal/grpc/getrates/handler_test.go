package getrates

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	"go.uber.org/zap/zaptest"
	pb "testTask/grpc/rates"
	"testTask/internal/usecase/actual_rate_get/domain"
)

func TestGetRates(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	logger := zaptest.NewLogger(t).Sugar()
	mockUsecase := NewMockusecase(ctrl)
	server := New(mockUsecase, logger)

	t.Run("successful GetRates", func(t *testing.T) {
		req := &pb.RateRequest{MarketID: "valid-market-id"}
		expectedResult := &domain.Result{
			Timestamp: 1633072800,
			AskPrice:  "1.23",
			BidPrice:  "1.20",
		}

		mockUsecase.EXPECT().
			Handle(context.Background(), req.MarketID).
			Return(expectedResult, nil)

		resp, err := server.GetRates(context.Background(), req)
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		if resp.Ask != expectedResult.AskPrice || resp.Bid != expectedResult.BidPrice {
			t.Fatalf("expected prices don't match, got %v/%v", resp.Ask, resp.Bid)
		}
	})

	t.Run("empty MarketID", func(t *testing.T) {
		req := &pb.RateRequest{MarketID: ""}

		resp, err := server.GetRates(context.Background(), req)
		if err == nil {
			t.Fatal("expected error, got none")
		}
		if resp != nil {
			t.Fatalf("expected nil response, got %v", resp)
		}
	})

	t.Run("error from usecase", func(t *testing.T) {
		req := &pb.RateRequest{MarketID: "valid-market-id"}
		expectedErr := errors.New("some error")

		mockUsecase.EXPECT().
			Handle(context.Background(), req.MarketID).
			Return(nil, expectedErr)

		resp, err := server.GetRates(context.Background(), req)
		if err == nil {
			t.Fatal("expected error, got none")
		}
		if resp != nil {
			t.Fatalf("expected nil response, got %v", resp)
		}
	})
}
