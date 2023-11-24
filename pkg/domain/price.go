package domain

type Price struct {
	Value int
}

func NewPrice(value int) (*Price, error) {
	if value <= 0 {
		return nil, NewInvalidArgumentError("price must be greater than 0")
	}
	return &Price{
		Value: value,
	}, nil
}
