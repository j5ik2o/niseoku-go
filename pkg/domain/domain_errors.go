package domain

// BaseError はエラーの基底構造体です。
type BaseError struct {
	Message string
}

// Error はエラーメッセージを返します。
func (e *BaseError) Error() string {
	return e.Message
}

// InvalidArgumentError は引数が不正な場合のエラーを表します。
type InvalidArgumentError struct {
	BaseError
}

// NewInvalidArgumentError は引数が不正な場合のエラーを生成します。
func NewInvalidArgumentError(message string) *InvalidArgumentError {
	return &InvalidArgumentError{
		BaseError{
			Message: message,
		},
	}
}
