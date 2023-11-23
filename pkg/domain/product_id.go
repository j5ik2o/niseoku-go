package domain

import "github.com/oklog/ulid/v2"

type ProductId struct {
	Value string
}

func NewProductId(value string) *ProductId {
	return &ProductId{
		Value: value,
	}
}

func GenerateProductId() *ProductId {
	ulid := ulid.Make()
	return &ProductId{
		Value: ulid.String(),
	}
}
