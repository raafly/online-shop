package repository

import (
	"context"
	"database/sql"

	"github.com/raafly/catering/model/domain"
)

type CustomerRepository interface {
	Register(ctx context.Context, tx *sql.Tx, customer domain.Customers) domain.Customers
	Login(ctx context.Context, db *sql.DB, customer domain.Customers) (domain.Customers, error)
	Update(ctx context.Context, tx *sql.Tx, customer domain.Customers) domain.Customers
	Delete(ctx context.Context, tx *sql.Tx, customerId int) 
	FindById(ctx context.Context, db *sql.DB, customerId int) (domain.Customers, error)
}