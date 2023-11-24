package domain

import "github.com/oklog/ulid/v2"

type AuctionId struct {
	Value string
}

func NewAuctionId(value string) (*AuctionId, error) {
	if value == "" {
		return nil, NewInvalidArgumentError("auction id must not be empty")
	}
	return &AuctionId{
		Value: value,
	}, nil
}

func GenerateAuctionId() *AuctionId {
	ulid := ulid.Make()
	return &AuctionId{
		Value: ulid.String(),
	}
}
