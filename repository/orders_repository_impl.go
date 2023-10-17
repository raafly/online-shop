package repository

import (
	"context"
	"database/sql"

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
