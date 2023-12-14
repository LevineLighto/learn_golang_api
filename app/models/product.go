package models

import "time"

type Product struct {
	Id        uint64    `json:"id"`
	Name      string    `json:"name"`
	SKU       string    `json:"sku"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
