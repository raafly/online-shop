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
	SQL := "INSERT INTO customers(username, email, password, address, no_telp) Values(?, ?, ?, ?, ?)"
	result, err := tx.ExecContext(ctx, SQL, customer.Username, customer.Email, customer.Password, customer.Address, customer.No_telp)
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

func (repository *CustomerRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, customer domain.Customers) domain.Customers {
	SQL := "UPDATE customers SET address = ? WHERE id = ?"
	_, err := tx.ExecContext(ctx, SQL, customer.Address, customer.Id)
	helper.PanicIfError(err)

	return customer
}

func (repository *CustomerRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, customerId int) {
	SQL := "DELETE from customers WHERE id = ?"
	_, err := tx.ExecContext(ctx, SQL, customerId)
	helper.PanicIfError(err)
}

func (repository *CustomerRepositoryImpl) FindById(ctx context.Context, db *sql.DB, customerId int) (domain.Customers, error) {
	SQL := "SELECT id, username, email, address, no_telp FROM customers WHERE id = ?"
	rows, err := db.QueryContext(ctx, SQL, customerId)
	helper.PanicIfError(err)
	defer rows.Close()

	customer := domain.Customers{}
	if rows.Next() {
		err := rows.Scan(&customer.Id, &customer.Username, &customer.Email, &customer.Address, &customer.No_telp)
		helper.PanicIfError(err)
		return customer, nil
	} else {
		return customer, errors.New("account not found")
	}
}