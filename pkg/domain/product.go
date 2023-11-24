package domain

type ProductType int

const (
	ProductTypeGeneric ProductType = iota
	ProductTypeDownloadSoftware
	ProductTypeCar
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

func NewProduct(id *ProductId, productType ProductType, name *ProductName, price *ProductPrice) *Product {
	return &Product{
		Id:          id,
		ProductType: productType,
		Name:        name,
		Price:       price,
		Status:      ProductStatusPrivate,
	}
}

func (p *Product) Publish() *Product {
	result := NewProduct(p.Id, p.ProductType, p.Name, p.Price)
	result.Status = ProductStatusPublic
	return result
}
