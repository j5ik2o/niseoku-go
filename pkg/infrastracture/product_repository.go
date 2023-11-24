package infrastracture

import "niseoku-go/pkg/domain"

type ProductRepository interface {
	Store(product *domain.Product) error
	FindById(productId *domain.ProductId) (*domain.Product, error)
	Delete(product *domain.Product) error
}
