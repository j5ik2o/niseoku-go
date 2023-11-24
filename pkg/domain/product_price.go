package domain

type ProductPrice struct {
	Value int
}

func NewProductPrice(value int) (*ProductPrice, error) {
	if value < 0 {
		return nil, NewInvalidArgumentError("product price must be greater than or equal to 0")
	}
	return &ProductPrice{
		Value: value,
	}, nil
}
