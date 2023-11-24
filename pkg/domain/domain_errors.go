package domain

type BaseError struct {
	Message string
}

func (e *BaseError) Error() string {
	return e.Message
}

type InvalidArgumentError struct {
	BaseError
}

func NewInvalidArgumentError(message string) *InvalidArgumentError {
	return &InvalidArgumentError{
		BaseError{
			Message: message,
		},
	}
}
