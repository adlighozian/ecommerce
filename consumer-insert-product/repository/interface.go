package repository

import "consumer-product-go/model"

type Product interface {
	CreateProduct(req []model.Product) error
}
