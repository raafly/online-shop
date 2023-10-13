package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/raafly/catering/helper"
	"github.com/raafly/catering/model/domain"
)

type ProductRepositoryImpl struct {
}

func NewProductRepository() *ProductRepositoryImpl {
	return &ProductRepositoryImpl{}
}

func (repository *ProductRepositoryImpl) GetAll(ctx context.Context, db *sql.DB) []domain.Products {
	SQL := "SELECT id, name, description, quantity, price, category FROM products"
	rows, err := db.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var products []domain.Products
	for rows.Next() {
		product := domain.Products{}
		err := rows.Scan(&product.Id, &product.Name, &product.Description, &product.Quantity, &product.Price, &product.Category)
		helper.PanicIfError(err)
		products = append(products, product)

	}
	return products
}

func (repository *ProductRepositoryImpl) Create(ctx context.Context, db *sql.DB, product domain.Products) domain.Products {
	SQL := "INSERT INTO products(id, name, description, quantity, price, category) VALUES(?, ?, ?, ?, ?, ?)"
	result, err := db.ExecContext(ctx, SQL, product.Id, product.Name, product.Description, product.Quantity, product.Price, product.Category)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	product.Id = int(id)
	return product 
}

func (repository *ProductRepositoryImpl) GetById(ctx context.Context, db *sql.DB, productId int) (domain.Products, error) {
	SQL := "SELECT id, name, description, quantity, price, category FROM products WHERE id = ?"
	rows, err := db.QueryContext(ctx, SQL, productId)
	helper.PanicIfError(err)
	defer rows.Close()

	product := domain.Products{}
	if rows.Next() {
		err := rows.Scan(&product.Id, &product.Name, &product.Description, &product.Quantity, &product.Price, &product.Category)
		helper.PanicIfError(err)
		return product, nil
	} else {
		return product, errors.New("product is not found")
	}
}

func (repository *ProductRepositoryImpl) Update(ctx context.Context, db *sql.DB, product domain.Products) domain.Products {
	SQL := "UPDATE products SET price = ?, quantity = ? WHERE id = ?"
	_, err := db.ExecContext(ctx, SQL, product.Price, product.Quantity ,product.Id)
	helper.PanicIfError(err)

	return product
}

func (repository *ProductRepositoryImpl) Delete(ctx context.Context, db *sql.DB, productId int) {
	SQL := "DELETE FROM products WHERE id = ?"
	_, err := db.ExecContext(ctx, SQL, productId)
	helper.PanicIfError(err)
}