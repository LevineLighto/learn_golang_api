package responses

import "learn_golang_api/app/models"

type ProductResponse struct {
	Id   uint64 `json:"id"`
	Name string `json:"name"`
	SKU  string `json:"sku"`
}

func NewProductResponse(product *models.Product) *ProductResponse {
	return &ProductResponse{
		Id:   product.Id,
		Name: product.Name,
		SKU:  product.SKU,
	}
}

func NewProductResponses(products []models.Product) []ProductResponse {

	var response []ProductResponse

	for _, product := range products {
		response = append(response, *NewProductResponse(&product))
	}

	return response

}
