package domain

import (
	"github.com/oklog/ulid/v2"
)

type ProductName struct {
	value string
}

func NewProductName(value string) (*ProductName, error) {
	if value == "" {
		return nil, NewInvalidArgumentError("product name must not be empty")
	}
	if len(value) > 255 {
		return nil, NewInvalidArgumentError("product name must not be longer than 255 characters")
	}
	return &ProductName{
		value: value,
	}, nil
}

func GenerateProductName() *ProductName {
	ulid := ulid.Make()
	return &ProductName{
		value: ulid.String(),
	}
}

func (p *ProductName) String() string {
	return p.value
}
