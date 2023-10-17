package service

import (
	"context"
	"database/sql"

	"github.com/go-playground/validator/v10"
	"github.com/raafly/catering/helper"
	"github.com/raafly/catering/model/domain"
	"github.com/raafly/catering/model/web"
	"github.com/raafly/catering/repository"
)

type OrderServiceImpl struct {
	OrderRepository		repository.OrderRepository
	DB 					*sql.DB
	Validate			*validator.Validate
}

func NewOrderServiceImpl(orderRepository repository.OrderRepository, DB *sql.DB, validate *validator.Validate) *OrderServiceImpl {
	return &OrderServiceImpl{
		OrderRepository: orderRepository,
		DB: DB,
		Validate: validate,
	}
}

func (service *OrderServiceImpl) Create(ctx context.Context, request web.OrderCreateRequest) web.OrderSuccess {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	defer helper.CommitOrRollback(tx)
	helper.PanicIfError(err)

	order := domain.Orders {
		Id_order: request.Id_order,
		Id_user: request.Id_user,
		Id_product: request.Id_product,		
		Note: request.Note,
	}

	order = service.OrderRepository.Create(ctx, tx, order)
	return helper.ToOrderReponse(order)
}