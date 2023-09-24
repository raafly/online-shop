package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/raafly/catering/helper"
	"github.com/raafly/catering/model/domain"
)

type CustomerRepositoryImpl struct {
}

func NewCustomerRepository() *CustomerRepositoryImpl {
	return &CustomerRepositoryImpl{}
}

func (repository *CustomerRepositoryImpl) Register(ctx context.Context, tx *sql.Tx, customer domain.Customers) domain.Customers {
	SQL := "INSERT INTO customers(username, email, password) Values(?, ?, ?)"
	result, err := tx.ExecContext(ctx, SQL, customer.Username, customer.Email, customer.Password)
	helper.PanicIfError(err) 

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	customer.Id = int(id)
	return customer 
}

func (repository *CustomerRepositoryImpl) Login(ctx context.Context, db *sql.DB, customer domain.Customers) (domain.Customers, error) {
	SQL := "SELECT email, password, role FROM customers WHERE email = ?"
	rows, err := db.QueryContext(ctx, SQL, customer.Email)
	helper.PanicIfError(err) 
	defer rows.Close()

	customer = domain.Customers{}
	if rows.Next() {
		err := rows.Scan(&customer.Email, &customer.Password, &customer.Role)
		helper.PanicIfError(err)
		return customer, nil
	} else {
		return customer, errors.New("user is not found")
	}
}
