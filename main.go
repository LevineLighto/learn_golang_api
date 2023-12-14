package main

import (
	"learn_golang_api/app"
	"learn_golang_api/app/controllers/product_controllers"
	"learn_golang_api/app/helpers"
	"learn_golang_api/app/middlewares"
	"learn_golang_api/app/repositories"
	"learn_golang_api/app/services"
	"net/http"

	"github.com/go-playground/validator/v10"
)

func NewServer(authMiddleware *middlewares.AuthMiddleware) *http.Server {
	return &http.Server{
		Addr:    "localhost:8000",
		Handler: authMiddleware,
	}
}

func InitializeServer() *http.Server {

	productRepository := repositories.NewProductRepository()
	db := app.NewDB()
	validate := validator.New()
	productService := services.NewProductService(productRepository, db, validate)
	productController := product_controllers.NewProductController(productService)
	router := app.NewRouter(productController)
	authMiddleware := middlewares.NewAuthMiddleware(router)
	server := NewServer(authMiddleware)

	return server

}

func main() {

	server := InitializeServer()
	err := server.ListenAndServe()
	helpers.PanicOnError(err)

}
