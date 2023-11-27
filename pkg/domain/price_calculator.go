package domain

import "fmt"

type PriceCalculator interface {
	CalculatePrice(auction *Auction, price *Price) (*Price, error)
	GetDelegate() *CalculatorDelegate
}

// ---

type ShippingFeeCalculator struct {
	delegate *CalculatorDelegate
}

func NewShippingFeeCalculator(next PriceCalculator) *ShippingFeeCalculator {
	var nextDelegate *CalculatorDelegate
	if next != nil {
		nextDelegate = next.GetDelegate()
	}
	delegate := NewCalculatorDelegate(
		"ShippingFeeCalculator",
		func(auction *Auction) (bool, bool) {
			return auction.product.productType == ProductTypeGeneric, true
		},
		func(auction *Auction, acc *Price) *Price {
			return acc.Add(NewPriceFromInt(BaseShippingFee))
		},
		nextDelegate,
	)
	return &ShippingFeeCalculator{delegate}
}

func (s *ShippingFeeCalculator) CalculatePrice(auction *Auction, acc *Price) (*Price, error) {
	return s.delegate.CalculatePrice(auction, acc)
}

func (s *ShippingFeeCalculator) GetDelegate() *CalculatorDelegate {
	return s.delegate
}

// ---

type CarShippingFeeCalculator struct {
	delegate *CalculatorDelegate
}

func NewCarShippingFeeCalculator(next PriceCalculator) *CarShippingFeeCalculator {
	var nextDelegate *CalculatorDelegate
	if next != nil {
		nextDelegate = next.GetDelegate()
	}
	delegate := NewCalculatorDelegate(
		"CarShippingFeeCalculator",
		func(auction *Auction) (bool, bool) {
			return auction.product.productType == ProductTypeCar, false
		},
		func(auction *Auction, acc *Price) *Price {
			fmt.Printf("CarShippingFeeCalculator: acc = %v\n", acc)
			return acc.Add(NewPriceFromInt(CarShippingFee))
		},
		nextDelegate,
	)
	return &CarShippingFeeCalculator{delegate}
}

func (c *CarShippingFeeCalculator) CalculatePrice(auction *Auction, price *Price) (*Price, error) {
	return c.delegate.CalculatePrice(auction, price)
}

func (c *CarShippingFeeCalculator) GetDelegate() *CalculatorDelegate {
	return c.delegate
}

// ---

type LuxuryCarTaxRateCalculator struct {
	delegate *CalculatorDelegate
}

func NewLuxuryCarTaxRateCalculator() *LuxuryCarTaxRateCalculator {
	delegate := NewCalculatorDelegate(
		"LuxuryCarTaxRateCalculator",
		func(auction *Auction) (bool, bool) {
			return auction.product.productType == ProductTypeCar && auction.highBidPrice.IsGreaterThanOrEqualTo(NewPriceFromInt(LuxuryCarPriceThreshold)), true
		},
		func(auction *Auction, acc *Price) *Price {
			return acc.Add(auction.highBidPrice.Multiply(LuxuryTaxRate))
		},
		nil,
	)
	return &LuxuryCarTaxRateCalculator{delegate}
}

func (c *LuxuryCarTaxRateCalculator) CalculatePrice(auction *Auction, acc *Price) (*Price, error) {
	return c.delegate.CalculatePrice(auction, acc)
}

func (c *LuxuryCarTaxRateCalculator) GetDelegate() *CalculatorDelegate {
	return c.delegate
}
