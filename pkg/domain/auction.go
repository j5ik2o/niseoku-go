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
	Id            *AuctionId
	Status        AuctionStatus
	Product       *Product
	StartDateTime *time.Time
	EndDateTime   *time.Time
	StartPrice    *Price
	SellerId      *UserAccountId
	HighBidderId  *UserAccountId
	HighBidPrice  *Price
	BuyerId       *UserAccountId
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
		Id:            id,
		Status:        AuctionStatusNotStarted,
		StartDateTime: startDateTime,
		EndDateTime:   endDateTime,
		Product:       product,
		StartPrice:    startPrice,
		SellerId:      sellerId,
	}, nil
}

func (a *Auction) Start(onStart EventHandler) *Auction {
	newAuction, _ := NewAuction(a.Id, a.Product, a.StartDateTime, a.EndDateTime, a.StartPrice, a.SellerId)
	newAuction.Status = AuctionStatusStarted
	onStart(newAuction)
	return newAuction
}

func (a *Auction) Close(onCloseWithNoBuyer EventHandler, onCloseWithBuyer EventHandler) *Auction {
	newAuction, _ := NewAuction(a.Id, a.Product, a.StartDateTime, a.EndDateTime, a.StartPrice, a.SellerId)
	newAuction.Status = AuctionStatusClosed
	if a.HighBidderId != nil {
		newAuction.BuyerId = a.HighBidderId
		// newAuction.BuyerPrice = ...
		onCloseWithBuyer(newAuction)
	} else {
		onCloseWithNoBuyer(newAuction)
	}
	return newAuction
}

func (a *Auction) Bid(price *Price, bidderId *UserAccountId) (*Auction, error) {
	if a.Status != AuctionStatusStarted {
		return nil, NewBidError("auction is not started")
	}
	if a.StartPrice.IsGreaterThan(price) {
		return nil, NewBidError("bid price must be higher than start price")
	}
	if a.HighBidPrice != nil && a.HighBidPrice.IsGreaterThanOrEqualTo(price) {
		return nil, NewBidError("bid price must be higher than high bid price")
	}
	result, _ := NewAuction(a.Id, a.Product, a.StartDateTime, a.EndDateTime, a.StartPrice, a.SellerId)
	result.HighBidPrice = price
	result.HighBidderId = bidderId
	return result, nil
}

func (a *Auction) GetSellerPrice() (*Price, error) {
	if a.HighBidPrice == nil {
		return nil, NewAuctionError("high bid price is not set")
	}
	hdp := float32(a.HighBidPrice.Value)
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
	if a.HighBidPrice == nil {
		return nil, NewAuctionError("high bid price is not set")
	}
	switch a.Product.ProductType {
	case ProductTypeGeneric:
		p, err := NewPrice(BaseShippingFee)
		if err != nil {
			return nil, err
		}
		return a.HighBidPrice.Add(p), nil
	case ProductTypeDownloadSoftware:
		return a.HighBidPrice, nil
	case ProductTypeCar:
		p, err := NewPrice(CarShippingFee)
		if err != nil {
			return nil, err
		}
		buyerPrice := a.HighBidPrice.Add(p)
		if a.HighBidPrice.IsGreaterThanOrEqualTo(&Price{LuxuryPriceThreshold}) {
			return buyerPrice.Add(a.HighBidPrice.Multiply(LuxuryTaxRate)), nil
		} else {
			return buyerPrice, nil
		}
	default:
		return nil, NewAuctionError("unknown product type")
	}
}
