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
	id          *ProductId
	productType ProductType
	name        *ProductName
	price       *ProductPrice
	status      ProductStatus
}

func NewProduct(id *ProductId, productType ProductType, name *ProductName, price *ProductPrice) *Product {
	return &Product{
		id:          id,
		productType: productType,
		name:        name,
		price:       price,
		status:      ProductStatusPrivate,
	}
}

func (p *Product) Publish() *Product {
	if p.status == ProductStatusPublic {
		return p
	}
	result := NewProduct(p.id, p.productType, p.name, p.price)
	result.status = ProductStatusPublic
	return result
}

func (p *Product) GetId() *ProductId {
	return p.id
}

func (p *Product) GetProductType() ProductType {
	return p.productType
}

func (p *Product) GetName() *ProductName {
	return p.name
}

func (p *Product) GetPrice() *ProductPrice {
	return p.price
}

func (p *Product) GetStatus() ProductStatus {
	return p.status
}

func (p *Product) IsPublic() bool {
	return p.status == ProductStatusPublic
}
