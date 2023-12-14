package product_requests

type CreateProductRequest struct {
	Name string `validate:"required,min=1,max=250" json:"name"`
	SKU  string `validate:"required,min=1,max=20" json:"sku"`
}
