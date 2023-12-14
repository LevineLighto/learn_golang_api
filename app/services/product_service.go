package services

import (
	"context"
	"learn_golang_api/app/requests/product_requests"
	"learn_golang_api/app/responses"
)

type ProductService interface {
	Create(ctx context.Context, request *product_requests.CreateProductRequest) *responses.ProductResponse
	Update(ctx context.Context, request *product_requests.UpdateProductRequest) *responses.ProductResponse
	Delete(ctx context.Context, productId uint64) bool
	Find(ctx context.Context, productId uint64) *responses.ProductResponse
	Get(ctx context.Context) []responses.ProductResponse
}
