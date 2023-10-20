package service

import (
	"context"

	"github.com/raafly/catering/model/web"
)

type OrderService interface {
	Create(ctx context.Context, request web.OrderCreateRequest) web.OrderSuccess
	GetById(ctx context.Context, orderId int) web.OrderDetailResponse
}