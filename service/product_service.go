package service

import (
	"context"

	"github.com/raafly/catering/model/web"
)

type ProductService interface {
	GetAll(ctx context.Context) []web.ProductResponse
	Create(ctx context.Context, request web.ProductsRequest) web.ProductResponse
	GetById(ctx context.Context, productId int) web.ProductResponse
	Update(ctx context.Context, request web.ProductUpdateRequest) web.ProductResponse
	Delete(ctx context.Context, productInt int)
}