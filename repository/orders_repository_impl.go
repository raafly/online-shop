package repository

import (
	"context"
	"database/sql"
	"errors"
	"log"

	"github.com/raafly/catering/helper"
	"github.com/raafly/catering/model/domain"
)

type OrderRepositoryImpl struct{
}

func NewOrderRepositoryImpl() *OrderRepositoryImpl {
	return &OrderRepositoryImpl{}
}

func (repository *OrderRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, order domain.Orders) domain.Orders {
	SQL := "INSERT INTO orders(id, id_user, id_product, note) VALUES(?, ?, ?, ?)"
	result, err := tx.ExecContext(ctx, SQL, order.Id_order, order.Id_user, order.Id_product, order.Note)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	order.Id_order = int(id)
	return order
}

func (repository *OrderRepositoryImpl) GetById(ctx context.Context, tx *sql.Tx, orderId int) (domain.Orders_detail, error) {
	SQL := "SELECT customers.id, orders.id, products.id, products.name, products.quantity, products.price, customers.address, customers.no_telp, orders.status FROM orders_detail JOIN customers ON orders_detail.id_user=customers.id JOIN products ON orders_detail.id_product=products.id JOIN orders ON orders_detail.id_order=orders.id WHERE id_order = ?"
	rows, err := tx.QueryContext(ctx, SQL, orderId)
	helper.PanicIfError(err)
	defer rows.Close()

	log.Println("repository",orderId)

	order := domain.Orders_detail{}
	if rows.Next() {
		err := rows.Scan(&order.Id_user, &order.Id_order, &order.Id_product, &order.Name, &order.Quantity, &order.Price, &order.Address, &order.No_telp, &order.Status)
		helper.PanicIfError(err)
		return order, nil
	} else {
		return order, errors.New("the order is not found")
	}
}

func (repository *OrderRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, order domain.Orders_detail) { 
	SQL := "UPDATE orders_detail SET quantity = ? WHERE id_order = ?"
	_, err := tx.ExecContext(ctx, SQL, order.Quantity, order.Id_order)
	helper.PanicIfError(err)
}

func (repository *OrderRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, orderId int) {
	SQL := "DELETE FROM orders_detail WHERE id_order = ?"
	_, err := tx.ExecContext(ctx, SQL, orderId)
	helper.PanicIfError(err)
}
