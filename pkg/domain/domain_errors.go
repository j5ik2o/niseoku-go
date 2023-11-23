package domain

type BaseError struct {
	Message string
}

func (e *BaseError) Error() string {
	return e.Message
}
