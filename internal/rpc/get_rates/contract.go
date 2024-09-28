//go:generate mockgen -package $GOPACKAGE -source $GOFILE -destination contract_mock.go
package get_rates

import (
	"context"
	"testTask/internal/usecase/actual_rate_get/domain"
)

type usecase interface {
	Handle(ctx context.Context, marketID string) (*domain.Result, error)
}
