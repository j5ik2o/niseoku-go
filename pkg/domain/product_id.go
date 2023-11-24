package domain

import "github.com/oklog/ulid/v2"

type ProductId struct {
	Value string
}

func NewProductId(value string) (*ProductId, error) {
	if value == "" {
		return nil, NewInvalidArgumentError("product id must not be empty")
	}
	return &ProductId{
		Value: value,
	}, nil
}

func GenerateProductId() *ProductId {
	ulid := ulid.Make()
	return &ProductId{
		Value: ulid.String(),
	}
}
