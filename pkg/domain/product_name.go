package domain

import "github.com/oklog/ulid/v2"

type ProductName struct {
	Value string
}

func NewProductName(value string) *ProductName {
	return &ProductName{
		Value: value,
	}
}

func GenerateProductName() *ProductName {
	ulid := ulid.Make()
	return &ProductName{
		Value: ulid.String(),
	}
}
