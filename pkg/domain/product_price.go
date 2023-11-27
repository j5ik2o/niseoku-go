package domain

// ProductPrice は商品の価格を表します。
type ProductPrice struct {
	value int
}

// NewProductPrice は商品の価格を生成します。
func NewProductPrice(value int) (*ProductPrice, error) {
	if value < 0 {
		return nil, NewInvalidArgumentError("product price must be greater than or equal to 0")
	}
	return &ProductPrice{
		value: value,
	}, nil
}

// GetValue は商品の価格を返します。
func (p *ProductPrice) GetValue() int {
	return p.value
}
