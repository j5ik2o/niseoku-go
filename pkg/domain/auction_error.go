package domain

// AuctionError はオークションに関するエラーを表します。
type AuctionError struct {
	BaseError
}

// NewAuctionError はオークションに関するエラーを生成します。
func NewAuctionError(message string) *AuctionError {
	return &AuctionError{
		BaseError{
			Message: message,
		},
	}
}

// BidError は入札に関するエラーを表します。
type BidError struct {
	BaseError
}

// NewBidError は入札に関するエラーを生成します。
//
// # 引数
// message: エラーメッセージ
//
// # 戻り値
// *BidError: 入札に関するエラー
func NewBidError(message string) *BidError {
	return &BidError{
		BaseError{
			Message: message,
		},
	}
}
