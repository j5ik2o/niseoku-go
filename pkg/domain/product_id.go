package domain

import "github.com/oklog/ulid/v2"

type ProductId struct {
	value string
}

func NewProductId(value string) (*ProductId, error) {
	if value == "" {
		return nil, NewInvalidArgumentError("product id must not be empty")
	}
	return &ProductId{
		value: value,
	}, nil
}

func GenerateProductId() *ProductId {
	ulid := ulid.Make()
	return &ProductId{
		value: ulid.String(),
	}
}

func (p *ProductId) String() string {
	return p.value
}
