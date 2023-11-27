package domain

import "github.com/oklog/ulid/v2"

// AuctionId はオークションのIDを表します。
type AuctionId struct {
	value string
}

func NewAuctionId(value string) (*AuctionId, error) {
	if value == "" {
		return nil, NewInvalidArgumentError("auction id must not be empty")
	}
	return &AuctionId{
		value: value,
	}, nil
}

func GenerateAuctionId() *AuctionId {
	ulid := ulid.Make()
	return &AuctionId{
		value: ulid.String(),
	}
}

func (a *AuctionId) String() string {
	return a.value
}

func (a *AuctionId) Equals(other *AuctionId) bool {
	return a.value == other.value
}
