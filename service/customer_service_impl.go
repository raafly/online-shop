package service

import (
	"context"
	"database/sql"

	"github.com/go-playground/validator/v10"
	"github.com/raafly/catering/exception"
	"github.com/raafly/catering/helper"
	"github.com/raafly/catering/model/domain"
	"github.com/raafly/catering/model/web"
	"github.com/raafly/catering/repository"

	"golang.org/x/crypto/bcrypt"
)

type CustomerServiceImpl struct {
	CustomerRepository repository.CustomerRepository
	DB *sql.DB
	Validate *validator.Validate
}

func NewCustomerService(customerRepository repository.CustomerRepository, DB *sql.DB, validate *validator.Validate) *CustomerServiceImpl {
	return &CustomerServiceImpl {
		CustomerRepository: customerRepository,
		DB: DB,
		Validate: validate,
	}
}

func (service *CustomerServiceImpl) Register(ctx context.Context ,request web.CustomerRegisterRequest) web.RegisterSuccess {
	// validasi 
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	// insert ke database
	tx, err := service.DB.Begin()
	defer helper.CommitOrRollback(tx)
	helper.PanicIfError(err)

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
    if err != nil {
        panic(err)
    }

	customer := domain.Customers {
		Username: request.Username,
		Email: request.Email,
		Password: string(hashedPassword),
	}

	customer = service.CustomerRepository.Register(ctx, tx, customer)

	return helper.ToCustomerResponse(customer)
	
}

func (service *CustomerServiceImpl) Login(ctx context.Context, request web.CustomerLoginRequest) web.LoginSuccess {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	customerData := domain.Customers {
		Email: request.Email,
		Password: request.Password,
	}

	customer, err := service.CustomerRepository.Login(ctx, service.DB, customerData)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	err = bcrypt.CompareHashAndPassword([]byte(customer.Password), []byte(customerData.Password))
	if err != nil {
		panic(exception.NewNotMatchError(err.Error()))
	}

	return helper.ToCustomerResponseLogin(customer)
}