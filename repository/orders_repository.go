package repository

import (
	"context"
	"database/sql"

	"github.com/raafly/catering/model/domain"
)

type OrderRepository interface {
	Create(ctx context.Context, tx *sql.Tx, order domain.Orders) domain.Orders
}