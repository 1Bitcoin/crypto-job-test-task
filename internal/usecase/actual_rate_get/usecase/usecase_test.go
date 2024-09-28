package usecase

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"testTask/internal/infrastructure/network"
	"testTask/internal/usecase/actual_rate_get/domain"
)

func TestHandle(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRateRepo := NewMockrateRepository(ctrl)
	mockClient := network.NewMockClient(ctrl)
	usecase := New(mockRateRepo, mockClient)

	marketID := "BTCUSD"
	url := GarantexURL + marketID

	t.Run("successful request", func(t *testing.T) {
		rate := domain.Rate{
			Timestamp: 1234567890,
			Asks:      []domain.Order{{Price: "1.23"}},
			Bids:      []domain.Order{{Price: "1.20"}},
		}
		rateJSON, _ := json.Marshal(rate)

		mockClient.EXPECT().Get(url).Return(&http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(bytes.NewReader(rateJSON)),
		}, nil)

		mockRateRepo.EXPECT().Save(gomock.Any(), gomock.Any()).Return(nil)

		result, err := usecase.Handle(context.Background(), marketID)

		assert.NoError(t, err)
		assert.Equal(t, rate.Timestamp, result.Timestamp)
		assert.Equal(t, rate.Asks[0].Price, result.AskPrice)
		assert.Equal(t, rate.Bids[0].Price, result.BidPrice)
	})

	t.Run("failed request to client", func(t *testing.T) {
		mockClient.EXPECT().Get(url).Return(nil, errors.New("client error"))

		_, err := usecase.Handle(context.Background(), marketID)

		assert.Error(t, err)
	})

	t.Run("failed to save rate", func(t *testing.T) {
		rate := domain.Rate{
			Timestamp: 1234567890,
			Asks:      []domain.Order{{Price: "1.23"}},
			Bids:      []domain.Order{{Price: "1.20"}},
		}
		rateJSON, _ := json.Marshal(rate)

		mockClient.EXPECT().Get(url).Return(&http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(bytes.NewReader(rateJSON)),
		}, nil)

		mockRateRepo.EXPECT().Save(gomock.Any(), gomock.Any()).Return(errors.New("save error"))

		_, err := usecase.Handle(context.Background(), marketID)

		assert.Error(t, err)
	})
}
