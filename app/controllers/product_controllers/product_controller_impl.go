package product_controllers

import (
	"learn_golang_api/app/helpers"
	"learn_golang_api/app/requests"
	"learn_golang_api/app/requests/product_requests"
	"learn_golang_api/app/responses"
	"learn_golang_api/app/services"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type ProductControllerImpl struct {
	ProductService services.ProductService
}

func NewProductController(productService services.ProductService) *ProductControllerImpl {
	return &ProductControllerImpl{
		ProductService: productService,
	}
}

func (controller *ProductControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	createProductRequest := product_requests.CreateProductRequest{}
	requests.Read(request, &createProductRequest)

	var apiResponse responses.APIResponse

	response := controller.ProductService.Create(request.Context(), &createProductRequest)
	if response.Id < 1 {
		apiResponse = responses.APIResponse{
			Message: "Failed to create product",
		}
		writer.WriteHeader(400)
	} else {
		apiResponse = responses.APIResponse{
			Message: "Successfully created a product",
			Data:    response,
		}
	}

	responses.Write(writer, apiResponse)

}

func (controller *ProductControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	updateProductRequest := product_requests.UpdateProductRequest{}
	requests.Read(request, &updateProductRequest)

	productId := params.ByName("productId")
	id, err := strconv.ParseUint(productId, 10, 64)
	helpers.PanicOnError(err)

	updateProductRequest.Id = id

	var apiResponse responses.APIResponse

	response := controller.ProductService.Update(request.Context(), &updateProductRequest)
	if response.Id < 1 {
		apiResponse = responses.APIResponse{
			Message: "Failed to update product",
		}
		writer.WriteHeader(400)
	} else {
		apiResponse = responses.APIResponse{
			Message: "Successfully updated the product",
			Data:    response,
		}
	}

	responses.Write(writer, apiResponse)

}

func (controller *ProductControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	productId := params.ByName("productId")
	id, err := strconv.ParseUint(productId, 10, 64)
	helpers.PanicOnError(err)

	var apiResponse responses.APIResponse

	response := controller.ProductService.Delete(request.Context(), id)
	if !response {
		apiResponse = responses.APIResponse{
			Message: "Failed to update product",
		}
		writer.WriteHeader(400)
	} else {
		apiResponse = responses.APIResponse{
			Message: "Successfully updated the product",
		}
	}

	responses.Write(writer, apiResponse)

}

func (controller *ProductControllerImpl) Find(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	productId := params.ByName("productId")
	id, err := strconv.ParseUint(productId, 10, 64)
	helpers.PanicOnError(err)

	var apiResponse responses.APIResponse

	response := controller.ProductService.Find(request.Context(), id)
	if response.Id < 0 {
		apiResponse = responses.APIResponse{
			Message: "Failed to fetch product",
		}
		writer.WriteHeader(404)
	} else {
		apiResponse = responses.APIResponse{
			Message: "Successfully fetched the product",
			Data:    response,
		}
	}

	responses.Write(writer, apiResponse)

}

func (controller *ProductControllerImpl) Get(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	var apiResponse responses.APIResponse

	response := controller.ProductService.Get(request.Context())

	apiResponse = responses.APIResponse{
		Message: "Successfully fetched products",
		Data:    response,
	}

	responses.Write(writer, apiResponse)

}
