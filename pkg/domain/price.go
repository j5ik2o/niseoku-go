package domain

type Price struct {
	Value int
}

func NewPrice(value int) (*Price, error) {
	if value <= 0 {
		return nil, NewInvalidArgumentError("price must be greater than 0")
	}
	return &Price{
		Value: value,
	}, nil
}

func (p *Price) IsGreaterThan(other *Price) bool {
	return p.Value > other.Value
}

func (p *Price) IsGreaterThanOrEqualTo(other *Price) bool {
	return p.Value >= other.Value
}

func (p *Price) IsLessThan(other *Price) bool {
	return p.Value < other.Value
}

func (p *Price) IsLessThanOrEqualTo(other *Price) bool {
	return p.Value <= other.Value
}

func (p *Price) IsEqualTo(other *Price) bool {
	return p.Value == other.Value
}
