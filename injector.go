//go:build wireinject
//+build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"github.com/raafly/catering/repository"
	"github.com/raafly/catering/config"
	"github.com/raafly/catering/route"
	"github.com/raafly/catering/middleware"
	"github.com/raafly/catering/service"
	"github.com/raafly/catering/controller"
	"github.com/go-playground/validator/v10"
)

var customerSet = wire.NewSet(
	repository.NewCustomerRepository,
	wire.Bind(new(repository.CustomerRepository), new(*repository.CustomerRepositoryImpl)),
	service.NewCustomerService,
	wire.Bind(new(service.CustomerService), new(*service.CustomerServiceImpl)),
	controller.NewCustomerController,
	wire.Bind(new(controller.CustomerController), new(*controller.CustomerControllerImpl)),

)

func InitializedEvent() *http.Server {
	wire.Build(
		config.NewDB,
		validator.New,
		route.NewRouter,
		customerSet,
		wire.Bind(new(http.Handler), new(*httprouter.Router)),
		middleware.NewAuthMiddleware,
		NewServer,
	)

	return nil
}