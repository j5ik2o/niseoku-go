package domain

import (
	"time"
)

type EventHandler func(auction *Auction)

type AuctionStatus int

const (
	AuctionStatusNotStarted AuctionStatus = iota
	AuctionStatusStarted
	AuctionStatusClosed
)

type Auction struct {
	id            *AuctionId
	status        AuctionStatus
	product       *Product
	startDateTime *time.Time
	endDateTime   *time.Time
	startPrice    *Price
	sellerId      *UserAccountId
	highBidderId  *UserAccountId
	highBidPrice  *Price
	buyerId       *UserAccountId
	buyerPrice    *Price
}

func (a *Auction) GetId() *AuctionId {
	return a.id
}

func (a *Auction) GetStatus() AuctionStatus {
	return a.status
}

func (a *Auction) GetProduct() *Product {
	return a.product
}

func (a *Auction) GetStartDateTime() *time.Time {
	return a.startDateTime
}

func (a *Auction) GetEndDateTime() *time.Time {
	return a.endDateTime
}

func (a *Auction) GetStartPrice() *Price {
	return a.startPrice
}

func (a *Auction) GetSellerId() *UserAccountId {
	return a.sellerId
}

func (a *Auction) GetHighBidderId() *UserAccountId {
	return a.highBidderId
}

func (a *Auction) GetHighBidPrice() *Price {
	return a.highBidPrice
}

func (a *Auction) GetBuyerId() *UserAccountId {
	return a.buyerId
}

func auctionNewNow() *time.Time {
	now := time.Now()
	return &now
}

func NewAuction(id *AuctionId, product *Product, startDateTime *time.Time, endDateTime *time.Time, startPrice *Price, sellerId *UserAccountId) (*Auction, error) {
	now := auctionNewNow()
	if startDateTime.Before(*now) {
		return nil, NewAuctionError("start date time must be future")
	}
	if endDateTime.Before(*startDateTime) {
		return nil, NewAuctionError("end date time must be after start date time")
	}
	return &Auction{
		id:            id,
		status:        AuctionStatusNotStarted,
		startDateTime: startDateTime,
		endDateTime:   endDateTime,
		product:       product,
		startPrice:    startPrice,
		sellerId:      sellerId,
	}, nil
}

func (a *Auction) Start(onStart EventHandler) *Auction {
	newAuction, _ := a.clone()
	newAuction.status = AuctionStatusStarted
	onStart(newAuction)
	return newAuction
}

func (a *Auction) clone() (*Auction, error) {
	return NewAuction(a.id, a.product, a.startDateTime, a.endDateTime, a.startPrice, a.sellerId)
}

func (a *Auction) Close(onCloseWithNoBuyer EventHandler, onCloseWithBuyer EventHandler) *Auction {
	newAuction, _ := a.clone()
	newAuction.status = AuctionStatusClosed
	if a.highBidderId != nil {
		newAuction.buyerId = a.highBidderId
		newAuction.buyerPrice = a.highBidPrice
		onCloseWithBuyer(newAuction)
	} else {
		onCloseWithNoBuyer(newAuction)
	}
	return newAuction
}

func (a *Auction) Bid(price *Price, bidderId *UserAccountId) (*Auction, error) {
	if a.status != AuctionStatusStarted {
		return nil, NewBidError("auction is not started")
	}
	if a.startPrice.IsGreaterThan(price) {
		return nil, NewBidError("bid price must be higher than start price")
	}
	if a.highBidPrice != nil && a.highBidPrice.IsGreaterThanOrEqualTo(price) {
		return nil, NewBidError("bid price must be higher than high bid price")
	}
	result, _ := a.clone()
	result.highBidPrice = price
	result.highBidderId = bidderId
	return result, nil
}

func (a *Auction) GetSellerPrice() (*Price, error) {
	if a.highBidPrice == nil {
		return nil, NewAuctionError("high bid price is not set")
	}
	hdp := float32(a.highBidPrice.GetValue())
	p, err := NewPrice(int(hdp - (hdp * 0.02)))
	if err != nil {
		return nil, err
	}
	return p, nil
}

const (
	BaseShippingFee      = 10
	CarShippingFee       = 1000
	LuxuryPriceThreshold = 50000
	LuxuryTaxRate        = 0.04
)

func (a *Auction) GetBuyerPrice() (*Price, error) {
	if a.highBidPrice == nil {
		return nil, NewAuctionError("high bid price is not set")
	}
	switch a.product.GetProductType() {
	case ProductTypeGeneric:
		p, err := NewPrice(BaseShippingFee)
		if err != nil {
			return nil, err
		}
		return a.highBidPrice.Add(p), nil
	case ProductTypeDownloadSoftware:
		return a.highBidPrice, nil
	case ProductTypeCar:
		p, err := NewPrice(CarShippingFee)
		if err != nil {
			return nil, err
		}
		buyerPrice := a.highBidPrice.Add(p)
		if a.highBidPrice.IsGreaterThanOrEqualTo(&Price{LuxuryPriceThreshold}) {
			return buyerPrice.Add(a.highBidPrice.Multiply(LuxuryTaxRate)), nil
		} else {
			return buyerPrice, nil
		}
	default:
		return nil, NewAuctionError("unknown product type")
	}
}
