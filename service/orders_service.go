package service

import (
	"context"

	"github.com/raafly/catering/model/web"
)

type OrderService interface {
	Create(ctx context.Context, request web.OrderCreateRequest) web.OrderSuccess
	Update(ctx context.Context, request web.OrderUpateRequest)
	Delete(ctx context.Context, orderId int)
	GetById(ctx context.Context, orderId int) web.OrderDetailResponse
}