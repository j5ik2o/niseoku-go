package domain

import (
	"github.com/oklog/ulid/v2"
)

type ProductName struct {
	Value string
}

func NewProductName(value string) (*ProductName, error) {
	if value == "" {
		return nil, NewInvalidArgumentError("product name must not be empty")
	}
	if len(value) > 255 {
		return nil, NewInvalidArgumentError("product name must not be longer than 255 characters")
	}
	return &ProductName{
		Value: value,
	}, nil
}

func GenerateProductName() *ProductName {
	ulid := ulid.Make()
	return &ProductName{
		Value: ulid.String(),
	}
}
