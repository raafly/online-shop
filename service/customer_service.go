package service

import (
	"context"
	"github.com/raafly/catering/model/web"
)

type CustomerService interface {
	Register(ctx context.Context, request web.CustomerRegisterRequest) web.RegisterSuccess
	Login(ctx context.Context, request web.CustomerLoginRequest) web.LoginSuccess
}