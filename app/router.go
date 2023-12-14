package app

import (
	"learn_golang_api/app/controllers/product_controllers"

	"github.com/julienschmidt/httprouter"
)

func NewRouter(productController product_controllers.ProductController) *httprouter.Router {

	router := httprouter.New()

	router.GET("/products", productController.Get)
	router.GET("/products/:productId", productController.Find)
	router.POST("/products", productController.Create)
	router.PUT("/products/:productId", productController.Update)
	router.DELETE("/products/:productId", productController.Delete)

	return router

}
