package domain

import "time"

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
		return nil, NewNewAuctionError("start date time must be future")
	}
	if endDateTime.Before(*startDateTime) {
		return nil, NewNewAuctionError("end date time must be after start date time")
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

func (a *Auction) Start(onStart ...EventHandler) *Auction {
	newAuction, _ := NewAuction(a.Id, a.Product, a.StartDateTime, a.EndDateTime, a.StartPrice, a.SellerId)
	newAuction.Status = AuctionStatusStarted
	for _, handler := range onStart {
		handler(newAuction)
	}
	return newAuction
}

func (a *Auction) Close(onClose ...EventHandler) *Auction {
	newAuction, _ := NewAuction(a.Id, a.Product, a.StartDateTime, a.EndDateTime, a.StartPrice, a.SellerId)
	newAuction.Status = AuctionStatusClosed
	for _, handler := range onClose {
		handler(newAuction)
	}
	return newAuction
}

func (a *Auction) Bid(price *Price, bidderId *UserAccountId) (*Auction, error) {
	if a.Status != AuctionStatusStarted {
		return nil, NewBidError("auction is not started")
	}
	if a.StartPrice.Value >= price.Value {
		return nil, NewBidError("bid price must be higher than start price")
	}
	if a.HighBidPrice != nil && a.HighBidPrice.Value >= price.Value {
		return nil, NewBidError("bid price must be higher than high bid price")
	}
	result, _ := NewAuction(a.Id, a.Product, a.StartDateTime, a.EndDateTime, a.StartPrice, a.SellerId)
	result.HighBidPrice = price
	result.HighBidderId = bidderId
	return result, nil
}
