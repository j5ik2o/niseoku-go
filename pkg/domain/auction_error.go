package domain

type NewAuctionError struct {
	BaseError
}

func NewNewAuctionError(message string) *NewAuctionError {
	return &NewAuctionError{
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
