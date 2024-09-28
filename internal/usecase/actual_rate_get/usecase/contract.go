//go:generate mockgen -package $GOPACKAGE -source $GOFILE -destination contract_mock.go
package usecase

import (
	"context"
	"testTask/internal/usecase/actual_rate_get/domain"
)

type rateRepository interface {
	Save(ctx context.Context, result domain.Result) error
}
