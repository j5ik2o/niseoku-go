package domain

type AuctionError struct {
	BaseError
}

func NewAuctionError(message string) *AuctionError {
	return &AuctionError{
		BaseError{
			Message: message,
		},
	}
}

type BidError struct {
	BaseError
}

func NewBidError(message string) *BidError {
	return &BidError{
		BaseError{
			Message: message,
		},
	}
}
