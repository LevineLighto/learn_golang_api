package services

import (
	"context"
	"database/sql"
	"learn_golang_api/app/helpers"
	"learn_golang_api/app/models"
	"learn_golang_api/app/repositories"
	"learn_golang_api/app/requests/product_requests"
	"learn_golang_api/app/responses"

	"github.com/go-playground/validator/v10"
)

type ProductServiceImpl struct {
	ProductRepository repositories.ProductRepository
	DB                *sql.DB
	Validate          *validator.Validate
}

func NewProductService(productRepository repositories.ProductRepository, DB *sql.DB, validate *validator.Validate) *ProductServiceImpl {
	return &ProductServiceImpl{
		ProductRepository: productRepository,
		DB:                DB,
		Validate:          validate,
	}
}

func (service *ProductServiceImpl) Create(ctx context.Context, request *product_requests.CreateProductRequest) *responses.ProductResponse {

	err := service.Validate.Struct(request)
	helpers.PanicOnError(err)

	transaction, err := service.DB.Begin()
	helpers.PanicOnError(err)
	defer helpers.CommitOrRollback(transaction)

	product := models.Product{
		Name: request.Name,
		SKU:  request.SKU,
	}

	if !service.ProductRepository.Save(ctx, transaction, &product) {
		return responses.NewProductResponse(&models.Product{})
	}

	return responses.NewProductResponse(&product)

}

func (service *ProductServiceImpl) Update(ctx context.Context, request *product_requests.UpdateProductRequest) *responses.ProductResponse {

	err := service.Validate.Struct(request)
	helpers.PanicOnError(err)

	transaction, err := service.DB.Begin()
	helpers.PanicOnError(err)
	defer helpers.CommitOrRollback(transaction)

	product := models.Product{
		Id:   request.Id,
		Name: request.Name,
		SKU:  request.SKU,
	}

	if !service.ProductRepository.Update(ctx, transaction, &product) {
		return responses.NewProductResponse(&models.Product{})
	}

	return responses.NewProductResponse(&product)

}

func (service *ProductServiceImpl) Delete(ctx context.Context, productId uint64) bool {

	transaction, err := service.DB.Begin()
	helpers.PanicOnError(err)
	defer helpers.CommitOrRollback(transaction)

	return service.ProductRepository.Delete(ctx, transaction, productId)

}

func (service *ProductServiceImpl) Find(ctx context.Context, productId uint64) *responses.ProductResponse {

	transaction, err := service.DB.Begin()
	helpers.PanicOnError(err)
	defer helpers.CommitOrRollback(transaction)

	product, err := service.ProductRepository.Find(ctx, transaction, productId)
	if err != nil {
		return responses.NewProductResponse(&models.Product{})
	}

	return responses.NewProductResponse(product)

}

func (service *ProductServiceImpl) Get(ctx context.Context) []responses.ProductResponse {

	transaction, err := service.DB.Begin()
	helpers.PanicOnError(err)
	defer helpers.CommitOrRollback(transaction)

	products, err := service.ProductRepository.Get(ctx, transaction)
	helpers.PanicOnError(err)

	return responses.NewProductResponses(products)

}
