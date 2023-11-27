package repository

import "niseoku-go/pkg/domain"

// ProductRepository は商品のリポジトリを表します。
type ProductRepository interface {
	Store(product *domain.Product) error
	FindById(productId *domain.ProductId) (*domain.Product, error)
	Delete(product *domain.Product) error
}
