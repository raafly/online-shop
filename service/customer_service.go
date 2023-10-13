package service

import (
	"context"
	"github.com/raafly/catering/model/web"
)

type CustomerService interface {
	Register(ctx context.Context, request web.CustomerRegisterRequest) web.RegisterSuccess
	Login(ctx context.Context, request web.CustomerLoginRequest) (response web.LoginSuccess, token string)
	Update(ctx context.Context, request web.CustomerUpdateRequest) web.RegisterSuccess
	Delete(ctx context.Context, customerId int) 
	FindById(ctx context.Context, customerId int) web.RegisterSuccess
}