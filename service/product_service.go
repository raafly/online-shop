package service

import (
	"context"
	"database/sql"

	"github.com/raafly/catering/model/web"
)

type ProducService interface {
	GetAll(ctx context.Context, db *sql.DB) []web.ProductResponse
	Create(ctx context.Context, request web.ProductsRequest) web.ProductResponse
	GetById(ctx context.Context, db *sql.DB, productId int) web.ProductResponse
	Update(ctx context.Context, db *sql.DB, request web.ProductsRequest) web.ProductResponse
	Delete(ctx context.Context, db *sql.DB, productName string)
}