package domain

import (
	"github.com/oklog/ulid/v2"
)

// ProductName は商品名を表します。
type ProductName struct {
	value string
}

// NewProductName は商品名を生成します。
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

// GenerateProductName は商品名を生成します。
func GenerateProductName() *ProductName {
	ulid := ulid.Make()
	return &ProductName{
		value: ulid.String(),
	}
}

// String は商品名を文字列で返します。
func (p *ProductName) String() string {
	return p.value
}
