package domain

type ProductType int

const (
	Generic ProductType = iota
	DownloadSoftware
	Car
)

type ProductStatus int

const (
	ProductStatusPublic ProductStatus = iota
	ProductStatusPrivate
)

type Product struct {
	Id          *ProductId
	ProductType ProductType
	Name        *ProductName
	Price       *ProductPrice
	Status      ProductStatus
}

func NewProduct(id *ProductId, name *ProductName, price *ProductPrice) *Product {
	return &Product{
		Id:     id,
		Name:   name,
		Price:  price,
		Status: ProductStatusPrivate,
	}
}

func (p *Product) Publish() *Product {
	result := NewProduct(p.Id, p.Name, p.Price)
	result.Status = ProductStatusPublic
	return result
}
