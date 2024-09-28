package repository

import (
	"context"
	"testTask/internal/usecase/actual_rate_get/domain"

	sq "github.com/Masterminds/squirrel"
	trmSqlx "github.com/avito-tech/go-transaction-manager/drivers/sqlx/v2"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type Repository struct {
	db        *sqlx.DB
	trmGetter *trmSqlx.CtxGetter
}

func New(db *sqlx.DB, trmGetter *trmSqlx.CtxGetter) *Repository {
	return &Repository{
		db:        db,
		trmGetter: trmGetter,
	}
}

func (r *Repository) Save(ctx context.Context, result domain.Result) error {
	queryBuilder := sq.StatementBuilder.PlaceholderFormat(sq.Dollar).
		Insert("rates").
		Columns(
			"timestamp",
			"ask_price",
			"bid_price",
		)

	queryBuilder = queryBuilder.Values(
		result.Timestamp,
		result.AskPrice,
		result.BidPrice,
	)

	query, args, err := queryBuilder.ToSql()
	if err != nil {
		return errors.Wrap(err, "build %s query")
	}
	_, err = r.trmGetter.DefaultTrOrDB(ctx, r.db).ExecContext(ctx, query, args...)
	if err != nil {
		return errors.Wrap(err, "unable to execute %s query")
	}

	return nil
}
