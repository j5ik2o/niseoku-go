package domain

type ProductPrice struct {
	Value int
}

func NewProductPrice(value int) *ProductPrice {
	return &ProductPrice{
		Value: value,
	}
}
