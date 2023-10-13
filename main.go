package main

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/raafly/catering/config"
	"github.com/raafly/catering/controller"
	"github.com/raafly/catering/helper"
	"github.com/raafly/catering/middleware"
	"github.com/raafly/catering/repository"
	"github.com/raafly/catering/service"
	"github.com/raafly/catering/route"
)

// func NewServer(authMiddleware *middleware.AuthMiddleware) *http.Server {
// 	return &http.Server{
// 		Addr:    "localhost:3000",
// 		Handler: authMiddleware,
// 	}
// }

func main() {

	db := config.NewDB()
	validate := validator.New()
	
	customerRepository := repository.NewCustomerRepository()
	customerService := service.NewCustomerService(customerRepository, db, validate)
	customerController := controller.NewCustomerController(customerService)

	productRepository := repository.NewProductRepository()
	productService := service.NewProductService(productRepository, db, validate)
	productController := controller.NewProductController(productService)

	router := route.NewRouter(customerController, productController)
	authMiddleware := middleware.NewAuthMiddleware(router)

	server := http.Server{
		Addr: "localhost:3000",
		Handler: authMiddleware,
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}