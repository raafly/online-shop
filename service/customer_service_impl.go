package service

import (
	"context"
	"database/sql"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"github.com/raafly/catering/config"
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
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

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
		Address: request.Address,
		No_telp: request.No_telp,
	}

	customer = service.CustomerRepository.Register(ctx, tx, customer)

	return helper.ToCustomerResponse(customer)
	
}

func (service *CustomerServiceImpl) Login(ctx context.Context, request web.CustomerLoginRequest) (response web.LoginSuccess, token string) {
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

	expTime := time.Now().Add(time.Minute * 1)
	claims := &config.JWTClaim{
		Username: customer.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer: "online-shop",
			ExpiresAt: jwt.NewNumericDate(expTime),
		},
	}

	tokenAlgo := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err = tokenAlgo.SignedString(config.JWT_KEY)
	helper.PanicIfError(err)

	return helper.ToCustomerResponseLogin(customer), token
}

func (service *CustomerServiceImpl) Update(ctx context.Context, request web.CustomerUpdateRequest) web.RegisterSuccess {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)	

	tx, err := service.DB.Begin()
	defer helper.CommitOrRollback(tx)
	helper.PanicIfError(err)

	customer, err := service.CustomerRepository.FindById(ctx, service.DB, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	customer = domain.Customers{
		Id: request.Id,
		Address: request.Address,
	}

	customer = service.CustomerRepository.Update(ctx, tx, customer)
	return helper.ToCustomerResponse(customer)	
}

func (service *CustomerServiceImpl) FindById(ctx context.Context, customerId int) web.RegisterSuccess {
	customer, err := service.CustomerRepository.FindById(ctx, service.DB, customerId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToCustomerResponse(customer)
}

func (service *CustomerServiceImpl) Delete(ctx context.Context, customerId int)  {
	customer, err := service.CustomerRepository.FindById(ctx, service.DB, customerId)
	if err != nil {
	   panic(exception.NewNotFoundError(err.Error()))
   }

   	tx, err := service.DB.Begin()
	defer helper.CommitOrRollback(tx)
	helper.PanicIfError(err)

 	service.CustomerRepository.Delete(ctx, tx, customer.Id)
}