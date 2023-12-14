package repositories

import (
	"context"
	"database/sql"
	"learn_golang_api/app/models"
)

type ProductRepository interface {
	Save(ctx context.Context, transaction *sql.Tx, product *models.Product) bool
	Update(ctx context.Context, transaction *sql.Tx, product *models.Product) bool
	Delete(ctx context.Context, transaction *sql.Tx, productId uint64) bool
	Find(ctx context.Context, transaction *sql.Tx, productId uint64) (*models.Product, error)
	Get(ctx context.Context, transaction *sql.Tx) ([]models.Product, error)
}
