package service

import (
	"context"
	"database/sql"

	"github.com/go-playground/validator/v10"
	"github.com/raafly/catering/helper"
	"github.com/raafly/catering/model/domain"
	"github.com/raafly/catering/model/web"
	"github.com/raafly/catering/repository"
	"github.com/raafly/catering/exception"
)

type ProductServiceImpl struct {
	ProductRepository repository.ProductRepository
	DB *sql.DB
	Validate *validator.Validate
}

func NewProductServiceImpl(productRepository repository.ProductRepository, DB *sql.DB, validate *validator.Validate) *ProductServiceImpl {
	return &ProductServiceImpl{
		ProductRepository: productRepository,
		DB: DB,
		Validate: validate,
	}
}

// func (service *ProductServiceImpl) GetAll(ctx context.Context, db *sql.DB) []web.ProductResponse {
// 	products := service.ProductRepository.GetAll(ctx, db)
// 	return helper.ToProductReponse(products)
// }
func (service *ProductServiceImpl) GetAll(ctx context.Context, db *sql.DB) []web.ProductResponse {
	products := service.ProductRepository.GetAll(ctx, db)
	return helper.ToProductResponses(products)
}

func (service *ProductServiceImpl) Create(ctx context.Context, request web.ProductsRequest) web.ProductResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	product := domain.Products {
		Name: request.Name,
		Description: request.Description,
		Quantity: request.Quantity,
		Price: request.Price,
	}	
	
	customer := service.ProductRepository.Create(ctx, service.DB, product)
	helper.PanicIfError(err)

	return helper.ToProductReponse(customer)
}

func (service *ProductServiceImpl) GetById(ctx context.Context, db *sql.DB, productId int) web.ProductResponse {
	product, err := service.ProductRepository.GetById(ctx, db, productId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	return helper.ToProductReponse(product)

}

func (service *ProductServiceImpl) Update(ctx context.Context, db *sql.DB ,request web.ProductsRequest) web.ProductResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	product, err := service.ProductRepository.GetById(ctx, db, request.Id)
	helper.PanicIfError(err)
	
	product, err = service.ProductRepository.Update(ctx, db, product)
	helper.PanicIfError(err)
	
	
	return helper.ToProductReponse(product)

}

func (service *ProductServiceImpl) Delete(ctx context.Context, db *sql.DB, productName string) {
	service.ProductRepository.Delete(ctx, db, productName)
}
