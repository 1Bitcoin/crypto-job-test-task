package usecase

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"net/http"
	"testTask/internal/infrastructure/network"
	"testTask/internal/usecase/actual_rate_get/domain"
)

const GarantexURL = "https://garantex.org/api/v2/depth?market="

type Usecase struct {
	rateRepository rateRepository
	client         network.Client
}

func New(rateRepository rateRepository, client network.Client) *Usecase {
	return &Usecase{
		rateRepository: rateRepository,
		client:         client,
	}
}

func (u *Usecase) Handle(ctx context.Context, marketID string) (*domain.Result, error) {
	// 1. build url
	url := fmt.Sprintf("%s%s", GarantexURL, marketID)

	// 2. Do http
	resp, err := u.client.Get(url)
	if err != nil {
		return nil, errors.Wrap(err, "do http request")
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.Errorf("wrong http response status code: %d", resp.StatusCode)
	}

	// 3. parse http response
	var rate domain.Rate
	if err = json.NewDecoder(resp.Body).Decode(&rate); err != nil {
		return nil, errors.Wrap(err, "unmarshal response")
	}

	if len(rate.Asks) == 0 || len(rate.Bids) == 0 {
		return nil, errors.New("no market data")
	}

	ask := rate.Asks[0]
	bid := rate.Bids[0]

	result := domain.Result{
		Timestamp: rate.Timestamp,
		AskPrice:  ask.Price,
		BidPrice:  bid.Price,
	}

	// 4. save result
	err = u.rateRepository.Save(ctx, result)
	if err != nil {
		return nil, errors.New("save to database")

	}

	return &result, nil
}
