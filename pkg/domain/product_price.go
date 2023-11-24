package domain

type ProductPrice struct {
	value int
}

func NewProductPrice(value int) (*ProductPrice, error) {
	if value < 0 {
		return nil, NewInvalidArgumentError("product price must be greater than or equal to 0")
	}
	return &ProductPrice{
		value: value,
	}, nil
}

func (p *ProductPrice) GetValue() int {
	return p.value
}
