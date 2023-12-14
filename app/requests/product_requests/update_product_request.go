package product_requests

type UpdateProductRequest struct {
	Id   uint64 `validate:"required" json:"id"`
	Name string `validate:"required,min=1,max=250" json:"name"`
	SKU  string `validate:"required,min=1,max=20" json:"sku"`
}
