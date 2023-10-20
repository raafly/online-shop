package test

import (
	"context"
	"fmt"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/raafly/catering/config"
	"github.com/raafly/catering/helper"
	"github.com/raafly/catering/repository"
	"github.com/raafly/catering/service"

	_ "github.com/go-sql-driver/mysql"
)

func TestRepositoryGetOrderById(t *testing.T) {
	db := config.NewDB()
	ctx := context.Background()

	tx, err := db.Begin()
	defer helper.CommitOrRollback(tx)
	helper.PanicIfError(err)

	_, err = repository.NewOrderRepositoryImpl().GetById(ctx, tx, 2)
	helper.PanicIfError(err)

}

func TestServiceGetOrderById(t *testing.T) {
	db := config.NewDB()
	ctx := context.Background()
	validate := validator.New()
	orderRepository := repository.OrderRepositoryImpl{}


	order := service.NewOrderServiceImpl(&orderRepository, db, validate).GetById(ctx, 2)	
	fmt.Println("data ->", order)
}