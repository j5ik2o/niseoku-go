package domain

// Price は金額を表します。
type Price struct {
	value int
}

// NewPriceFromInt は int 型の値から Price を生成します。
func NewPriceFromInt(value int) *Price {
	if value <= 0 {
		panic("price must be greater than 0")
	}
	return &Price{
		value: value,
	}
}

// NewPrice は金額を生成します。
func NewPrice(value int) (*Price, error) {
	if value <= 0 {
		return nil, NewInvalidArgumentError("price must be greater than 0")
	}
	return &Price{
		value: value,
	}, nil
}

// Add は金額を加算します。
func (p *Price) Add(other *Price) *Price {
	return NewPriceFromInt(p.value + other.value)
}

// Subtract は金額を減算します。
func (p *Price) Subtract(other *Price) *Price {
	return NewPriceFromInt(p.value - other.value)
}

// Multiply は金額を乗算します。
func (p *Price) Multiply(rate float32) *Price {
	return NewPriceFromInt(int(float32(p.value) * rate))
}

// IsGreaterThan は金額が指定した金額より大きいかどうかを判定します。
func (p *Price) IsGreaterThan(other *Price) bool {
	return p.value > other.value
}

// IsGreaterThanOrEqualTo は金額が指定した金額以上かどうかを判定します。
func (p *Price) IsGreaterThanOrEqualTo(other *Price) bool {
	return p.value >= other.value
}

// IsLessThan は金額が指定した金額より小さいかどうかを判定します。
func (p *Price) IsLessThan(other *Price) bool {
	return p.value < other.value
}

// IsLessThanOrEqualTo は金額が指定した金額以下かどうかを判定します。
func (p *Price) IsEqualTo(other *Price) bool {
	return p.value == other.value
}

// GetValue は金額の値を取得します。
func (p *Price) GetValue() int {
	return p.value
}
