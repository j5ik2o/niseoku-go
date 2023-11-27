package domain

type Resolver func(auction *Auction) (bool, bool)
type Calculate func(auction *Auction, acc *Price) *Price

type CalculatorDelegate struct {
	name      string
	resolver  Resolver
	calculate Calculate
	next      *CalculatorDelegate
}

func NewCalculatorDelegate(name string, resolver Resolver, calculate Calculate, next *CalculatorDelegate) *CalculatorDelegate {
	return &CalculatorDelegate{
		name,
		resolver,
		calculate,
		next,
	}
}

func (s *CalculatorDelegate) CalculatePrice(auction *Auction, currentPrice *Price) (*Price, error) {
	condition, controlBreak := s.resolver(auction)
	if condition {
		price := s.calculate(auction, currentPrice)
		if controlBreak {
			return price, nil
		} else if s.next != nil {
			return s.next.CalculatePrice(auction, price)
		} else {
			return price, nil
		}
	} else if s.next != nil {
		return s.next.CalculatePrice(auction, currentPrice)
	} else {
		return currentPrice, nil
	}
}
