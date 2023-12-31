package test

import (
	"context"
	"database/sql"
	"fmt"
	_ "fmt"
	"testing"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/raafly/catering/helper"
	"github.com/raafly/catering/model/domain"
	"github.com/raafly/catering/repository"
)

func NewDB() *sql.DB {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/catering")

	helper.PanicIfError(err)

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(5 *time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	return db
}

// func SetupRouter() http.Handler {
// 	db := ConnDB()
// 	validate := validator.New()
// 	productRepository := repository.NewProducRepository()
// 	customerService := service.NewCustomerService(productRepository, db, validate)
// 	customerController := controller.NewCustomerController(customerService)
// 	router := route.NewRouter(customerController)

// 	return router
// }

// func TestRegister(t *testing.T) {

// 	router :=  SetupRouter()

// 	requestBody := strings.NewReader(`{
// 		"username" : "rafly",
// 		"email" : "rafra123@gmail.com"
// 		"password" : "raafly",
// 	}`)
// 	request := httptest.NewRequest(http.MethodPost, "http://localhost:300/api/register", requestBody)
// 	request.Header.Add("Content-Type", "application/json")
// 	recorder := httptest.NewRecorder()

// 	router.ServeHTTP(recorder, request)

// 	response := recorder.Result()
// 	body, err := io.ReadAll(response.Body)
// 	defer response.Body.Close()
// 	helper.PanicIfError(err)

// 	fmt.Println(string(body))

// }

// func TestQueryCreateProduct(t *testing.T) {
// 	ctx := context.Background()
// 	db := NewDB()

// 	product := domain.Products {
// 		Name: "tahu",
// 		Description: "tahu goreng",
// 		Quantity: 70,
// 		Price: 9.000,
// 	}

// 	products := repository.NewProducRepository().Create(ctx, db, product)
// 	fmt.Println(products)
// }

// func TestGetQueryProduct(t *testing.T) {
// 	ctx := context.Background()
// 	db := NewDB()

// 	products, err := repository.NewProducRepository().GetById(ctx, db, 1)
// 	helper.PanicIfError(err)
// 	fmt.Println(products)
// }

// func TestUpdateQueryProduct(t *testing.T) {
// 	ctx := context.Background()
// 	db := NewDB()

// 	product := domain.Products {
// 		Id: 1,
// 		Price: 20.000,
// 	}

// 	products := repository.NewProducRepository().Update(ctx, db, product)
// 	fmt.Println(products)
// }

// func TestGetAllQueryProduct(t *testing.T) {
// 	ctx := context.Background()
// 	db := NewDB()

// 	products := repository.NewProducRepository().GetAll(ctx, db)
// 	fmt.Println(products)
// }

// func TestDeleteQueryProduct(t *testing.T) {
// 	ctx := context.Background()
// 	db := NewDB()

// 	repository.NewProducRepository().Delete(ctx, db, "tempe")
// }

func TestUp(t *testing.T) {
	db := NewDB()
	ctx := context.Background()
	product := domain.Products{
		Id: 5,
		Price: 100.0000,
		Quantity: 70,
	}

	product = repository.NewProductRepository().Update(ctx, db, product)
	fmt.Println(product)

}