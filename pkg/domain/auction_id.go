package domain

import "github.com/oklog/ulid/v2"

type AuctionId struct {
	Value string
}

func NewAuctionId(value string) *AuctionId {
	return &AuctionId{
		Value: value,
	}
}

func GenerateAuctionId() *AuctionId {
	ulid := ulid.Make()
	return &AuctionId{
		Value: ulid.String(),
	}
}
