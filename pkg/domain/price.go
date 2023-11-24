package domain

type Price struct {
	value int
}

func NewPriceFromInt(value int) *Price {
	return &Price{
		value: value,
	}
}

func NewPrice(value int) (*Price, error) {
	if value <= 0 {
		return nil, NewInvalidArgumentError("price must be greater than 0")
	}
	return &Price{
		value: value,
	}, nil
}

func (p *Price) Add(other *Price) *Price {
	return &Price{
		value: p.value + other.value,
	}
}

func (p *Price) Subtract(other *Price) *Price {
	return &Price{
		value: p.value - other.value,
	}
}

func (p *Price) Multiply(rate float32) *Price {
	return &Price{
		value: int(float32(p.value) * rate),
	}
}

func (p *Price) IsGreaterThan(other *Price) bool {
	return p.value > other.value
}

func (p *Price) IsGreaterThanOrEqualTo(other *Price) bool {
	return p.value >= other.value
}

func (p *Price) IsLessThan(other *Price) bool {
	return p.value < other.value
}

func (p *Price) IsEqualTo(other *Price) bool {
	return p.value == other.value
}

func (p *Price) GetValue() int {
	return p.value
}
