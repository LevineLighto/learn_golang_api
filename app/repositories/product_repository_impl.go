package repositories

import (
	"context"
	"database/sql"
	"errors"
	"learn_golang_api/app/helpers"
	"learn_golang_api/app/models"

	_ "github.com/jackc/pgx/v5/stdlib"
)

type ProductRepositoryImpl struct{}

func NewProductRepository() *ProductRepositoryImpl {
	return &ProductRepositoryImpl{}
}

func (repository *ProductRepositoryImpl) Save(ctx context.Context, transaction *sql.Tx, product *models.Product) bool {

	query := "INSERT INTO public.products (name, sku) values ($1, $2) RETURNING id, name, sku, created_at, updated_at"
	row := transaction.QueryRowContext(ctx, query, product.Name, product.SKU)

	err := row.Scan(&product.Id, &product.Name, &product.SKU, &product.CreatedAt, &product.UpdatedAt)
	helpers.PanicOnError(err)

	return product.Id >= 1

}

func (repository *ProductRepositoryImpl) Update(ctx context.Context, transaction *sql.Tx, product *models.Product) bool {

	query := "UPDATE public.products SET name = $1, sku = $2, updated_at = CURRENT_TIMESTAMP WHERE id = $3"
	result, err := transaction.ExecContext(ctx, query, product.Name, product.SKU, product.Id)
	helpers.PanicOnError(err)

	if row, err := result.RowsAffected(); row < 1 || err != nil {
		return false
	}

	return true

}

func (repository *ProductRepositoryImpl) Delete(ctx context.Context, transaction *sql.Tx, productId uint64) bool {

	query := "DELETE FROM public.products WHERE id = $1"
	result, err := transaction.ExecContext(ctx, query, productId)
	helpers.PanicOnError(err)

	if row, err := result.RowsAffected(); row < 1 || err != nil {
		return false
	}

	return true

}

func (repository *ProductRepositoryImpl) Find(ctx context.Context, transaction *sql.Tx, productId uint64) (*models.Product, error) {

	query := "SELECT id, name, sku, created_at, updated_at FROM public.products WHERE id = $1"
	rows, err := transaction.QueryContext(ctx, query, productId)
	helpers.PanicOnError(err)
	defer rows.Close()

	product := models.Product{}
	if rows.Next() {

		err := rows.Scan(&product.Id, &product.Name, &product.SKU, &product.CreatedAt, &product.UpdatedAt)
		helpers.PanicOnError(err)

		return &product, nil

	}

	return &product, errors.New("could not find product")

}

func (repository *ProductRepositoryImpl) Get(ctx context.Context, transaction *sql.Tx) ([]models.Product, error) {

	query := "SELECT id, name, sku FROM public.products ORDER BY created_at DESC"
	rows, err := transaction.QueryContext(ctx, query)
	helpers.PanicOnError(err)
	defer rows.Close()

	var products []models.Product

	for rows.Next() {

		product := models.Product{}
		err := rows.Scan(&product.Id, &product.Name, &product.SKU)
		helpers.PanicOnError(err)

		products = append(products, product)

	}

	return products, nil

}
