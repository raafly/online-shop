package test

import (
	"context"
	"database/sql"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/raafly/catering/controller"
	"github.com/raafly/catering/helper"
	"github.com/raafly/catering/model/domain"
	"github.com/raafly/catering/repository"
	"github.com/raafly/catering/route"
	"github.com/raafly/catering/service"
	"github.com/go-playground/validator/v10"
)

func ConnDB() *sql.DB {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/catering")

	helper.PanicIfError(err)

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(5 *time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	return db
}

func SetupRouter() http.Handler {
	db := ConnDB()
	validate := validator.New()
	customerRepository := repository.NewCustomerRepository()
	customerService := service.NewCustomerService(customerRepository, db, validate)
	customerController := controller.NewCustomerController(customerService)
	router := route.NewRouter(customerController)

	return router
}

func TestRegister(t *testing.T) {

	router :=  SetupRouter()

	requestBody := strings.NewReader(`{
		"username" : "rafly",
		"email" : "rafra123@gmail.com"
		"password" : "raafly",
	}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:300/api/register", requestBody)
	request.Header.Add("Content-Type", "application/json")
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	body, err := io.ReadAll(response.Body)
	defer response.Body.Close()
	helper.PanicIfError(err)

	fmt.Println(string(body))

}

func TestQuery(t *testing.T) {
	ctx := context.Background()
	db := ConnDB()

	customer := domain.Customers {
		Email: "aku123@gmail.com",
		Password: "taku123",
	}

	login, err := repository.NewCustomerRepository().Login(ctx, db, customer)
	helper.PanicIfError(err)
	fmt.Println(login)
}