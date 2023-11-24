package memory

import "niseoku-go/pkg/domain"

type ProductRepositoryInMemory struct {
	products map[*domain.ProductId]*domain.Product
}

func NewProductRepositoryInMemory() *ProductRepositoryInMemory {
	return &ProductRepositoryInMemory{
		products: make(map[*domain.ProductId]*domain.Product),
	}
}

func (r *ProductRepositoryInMemory) Store(product *domain.Product) error {
	r.products[product.Id] = product
	return nil
}

func (r *ProductRepositoryInMemory) FindById(productId *domain.ProductId) (*domain.Product, error) {
	return r.products[productId], nil
}

func (r *ProductRepositoryInMemory) Delete(product *domain.Product) error {
	delete(r.products, product.Id)
	return nil
}
