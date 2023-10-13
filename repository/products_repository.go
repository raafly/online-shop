package repository

import (
	"context"
	"database/sql"
	"github.com/raafly/catering/model/domain"
)

type ProductRepository interface {
	GetAll(ctx context.Context, db *sql.DB) []domain.Products
	Create(ctx context.Context, db *sql.DB, product domain.Products) domain.Products
	GetById(ctx context.Context, db *sql.DB, productId int) (domain.Products, error)
	Update(ctx context.Context, db *sql.DB, product domain.Products) domain.Products
	Delete(ctx context.Context, db *sql.DB, productId int) 
}