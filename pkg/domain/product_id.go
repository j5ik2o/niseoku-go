package domain

import "github.com/oklog/ulid/v2"

// ProductId は商品のIDを表します。
type ProductId struct {
	value string
}

// NewProductId は商品のIDを生成します。
func NewProductId(value string) (*ProductId, error) {
	if value == "" {
		return nil, NewInvalidArgumentError("product id must not be empty")
	}
	return &ProductId{
		value: value,
	}, nil
}

// GenerateProductId は商品のIDを生成します。
func GenerateProductId() *ProductId {
	ulid := ulid.Make()
	return &ProductId{
		value: ulid.String(),
	}
}

// String は商品のIDを文字列で返します。
func (p *ProductId) String() string {
	return p.value
}
