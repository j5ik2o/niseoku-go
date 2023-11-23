package domain

type PriceError struct {
	BaseError
}

func NewPriceError(message string) *PriceError {
	return &PriceError{
		BaseError{
			Message: message,
		},
	}
}
