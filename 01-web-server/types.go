package main

import "math/rand"

type CreateProductRequest struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type UpdateProductRequest struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type Product struct {
	ID        int64   `json:"id"`
	Name      string  `json:"name"`
	Price     float64 `json:"price"`
	Currency  string  `json:"currency"`
	InStock   bool    `json:"in_stock"`
	IsDeleted bool    `json:"is_deleted"`
}

func NewProduct(name string, price float64) *Product {
	return &Product{
		ID:       int64(rand.Int()),
		Name:     name,
		Price:    price,
		Currency: "euro",
	}
}
