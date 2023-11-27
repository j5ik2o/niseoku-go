package domain

import (
	"time"
)

// EventHandler はイベントハンドラを表します。
type EventHandler func(auction *Auction)

// AuctionStatus はオークションの状態を表します。
type AuctionStatus int

const (
	// AuctionStatusNotStarted は未開始を表します。
	AuctionStatusNotStarted AuctionStatus = iota
	// AuctionStatusStarted は開始済みを表します。
	AuctionStatusStarted
	// AuctionStatusClosed は終了済みを表します。
	AuctionStatusClosed
)

const (
	// BaseShippingFee は基本送料を表します。
	BaseShippingFee = 10
	// CarShippingFee は車の送料を表します。
	CarShippingFee = 1000
	// LuxuryCarPriceThreshold は高級車の価格閾値を表します。
	LuxuryCarPriceThreshold = 50000
	// LuxuryTaxRate は贅沢品の税率を表します。
	LuxuryTaxRate = 0.04
)

// Auction はオークション集約を表します。
type Auction struct {
	id            *AuctionId
	status        AuctionStatus
	product       *Product
	startDateTime *time.Time
	endDateTime   *time.Time
	startPrice    *Price
	sellerId      *UserAccountId
	highBidderId  *UserAccountId
	highBidPrice  *Price
	buyerId       *UserAccountId
	buyerPrice    *Price
	clock         Clock
}

func (a *Auction) SetClock(clock Clock) {
	a.clock = clock
}

func (a *Auction) GetId() *AuctionId {
	return a.id
}

func (a *Auction) GetStatus() AuctionStatus {
	return a.status
}

func (a *Auction) GetProduct() *Product {
	return a.product
}

func (a *Auction) GetStartDateTime() *time.Time {
	return a.startDateTime
}

func (a *Auction) GetEndDateTime() *time.Time {
	return a.endDateTime
}

func (a *Auction) GetStartPrice() *Price {
	return a.startPrice
}

func (a *Auction) GetSellerId() *UserAccountId {
	return a.sellerId
}

func (a *Auction) GetHighBidderId() *UserAccountId {
	return a.highBidderId
}

func (a *Auction) GetHighBidPrice() *Price {
	return a.highBidPrice
}

func (a *Auction) GetBuyerId() *UserAccountId {
	return a.buyerId
}

// clone はオークション集約を複製する関数
func (a *Auction) clone() (*Auction, error) {
	return newAuctionWithNoValidation(a.clock, a.id, a.product, a.startDateTime, a.endDateTime, a.startPrice, a.sellerId)
}

func newAuctionWithNoValidation(clock Clock, id *AuctionId, product *Product, startDateTime *time.Time, endDateTime *time.Time, startPrice *Price, sellerId *UserAccountId) (*Auction, error) {
	return &Auction{
		id:            id,
		status:        AuctionStatusNotStarted,
		startDateTime: startDateTime,
		endDateTime:   endDateTime,
		product:       product,
		startPrice:    startPrice,
		sellerId:      sellerId,
		clock:         clock,
	}, nil
}

var (
	buyerPriceCalculator = NewShippingFeeCalculator(NewCarShippingFeeCalculator(NewLuxuryCarTaxRateCalculator()))
)

// NewAuction はオークション集約を生成する関数
//
// オークション集約を生成する際には、以下のルールに従う必要がある
// - 開始日時は未来であること
// - 終了日時は開始日時より後であること
// - 開始価格は0より大きいこと
//
// # 引数
// - clock: クロック
// - id: オークションID
// - product: 商品
// - startDateTime: 開始日時
// - endDateTime: 終了日時
// - startPrice: 開始価格
// - sellerId: 出品者ID
//
// # 戻り値
// - Auction オークション集約
// - エラー
func NewAuction(clock Clock, id *AuctionId, product *Product, startDateTime *time.Time, endDateTime *time.Time, startPrice *Price, sellerId *UserAccountId) (*Auction, error) {
	now := clock.Now()
	if !product.IsPublic() {
		return nil, NewAuctionError("product is private")
	}
	if startDateTime.Before(*now) {
		return nil, NewAuctionError("start date time must be future")
	}
	if endDateTime.Before(*startDateTime) {
		return nil, NewAuctionError("end date time must be after start date time")
	}
	return &Auction{
		id:            id,
		status:        AuctionStatusNotStarted,
		startDateTime: startDateTime,
		endDateTime:   endDateTime,
		product:       product,
		startPrice:    startPrice,
		sellerId:      sellerId,
		clock:         &SystemClock{},
	}, nil
}

// Start はオークションを開始する関数
//
// オークションを開始する際には、以下のルールに従う必要がある
// - オークションは開始前であること
//
// # 引数
// - onStart: オークション開始時のイベントハンドラ
//
// # 戻り値
// - オークション集約
func (a *Auction) Start(onStart EventHandler) (*Auction, error) {
	if a.status != AuctionStatusNotStarted {
		return nil, NewAuctionError("auction is not not started")
	}
	if a.startDateTime.After(*a.clock.Now()) {
		return nil, NewAuctionError("auction start date time is not after now")
	}
	newAuction, _ := a.clone()
	newAuction.status = AuctionStatusStarted
	onStart(newAuction)
	return newAuction, nil
}

// Close はオークションを終了する関数
//
// オークションを終了する際には、以下のルールに従う必要がある
// - オークションは開始済みであること
//
// # 引数
// - onCloseWithNoBuyer: 落札者がいなかった場合のイベントハンドラ
// - onCloseWithBuyer: 落札者がいた場合のイベントハンドラ
//
// # 戻り値
// - Auction オークション集約
func (a *Auction) Close(onCloseWithNoBuyer EventHandler, onCloseWithBuyer EventHandler) *Auction {
	newAuction, _ := a.clone()
	newAuction.status = AuctionStatusClosed
	if a.highBidderId != nil {
		newAuction.buyerId = a.highBidderId
		newAuction.buyerPrice = a.highBidPrice
		onCloseWithBuyer(newAuction)
	} else {
		onCloseWithNoBuyer(newAuction)
	}
	return newAuction
}

// Bid は入札する関数
//
// 入札する際には、以下のルールに従う必要がある
// - オークションは開始済みであること
// - 入札価格は開始価格より高いこと
// - 入札価格は現在の最高入札価格より高いこと
//
// # 引数
// - price: 入札価格
// - bidderId: 入札者ID
//
// # 戻り値
// - Auction オークション集約
// - エラー
func (a *Auction) Bid(price *Price, bidderId *UserAccountId) (*Auction, error) {
	if a.status != AuctionStatusStarted {
		return nil, NewBidError("auction is not started")
	}
	if a.startPrice.IsGreaterThan(price) {
		return nil, NewBidError("bid price must be higher than start price")
	}
	if a.highBidPrice != nil && a.highBidPrice.IsGreaterThanOrEqualTo(price) {
		return nil, NewBidError("bid price must be higher than high bid price")
	}
	result, _ := a.clone()
	result.highBidPrice = price
	result.highBidderId = bidderId
	return result, nil
}

// CancelBid は入札をキャンセルする関数
//
// 入札をキャンセルする際には、以下のルールに従う必要がある
// - オークションは開始済みであること
// - 入札者は最高入札者であること
//
// # 引数
// - bidderId: 入札者ID
//
// # 戻り値
// - Auction オークション集約
// - エラー
func (a *Auction) CancelBid(bidderId *UserAccountId) (*Auction, error) {
	if a.status != AuctionStatusStarted {
		return nil, NewBidError("auction is not started")
	}
	if a.highBidderId == nil {
		return nil, NewBidError("high bidder is not set")
	}
	if !a.highBidderId.Equals(bidderId) {
		return nil, NewBidError("bidder is not high bidder")
	}
	result, _ := a.clone()
	result.highBidPrice = nil
	result.highBidderId = nil
	return result, nil
}

// GetSellerPrice は出品者が受け取る金額を取得する関数
//
// 出品者が受け取る金額を取得する際には、以下のルールに従う必要がある
// - オークションは終了済みであること
// - 最高入札価格が設定されていること
//
// # 戻り値
// - Price 出品者が受け取る金額
// - エラー
func (a *Auction) GetSellerPrice() (*Price, error) {
	if a.highBidPrice == nil {
		return nil, NewAuctionError("high bid price is not set")
	}
	hdp := float32(a.highBidPrice.GetValue())
	p, err := NewPrice(int(hdp - (hdp * 0.02)))
	if err != nil {
		return nil, err
	}
	return p, nil
}

// GetBuyerPrice は落札者が支払う金額を取得する関数
//
// 落札者が支払う金額を取得する際には、以下のルールに従う必要がある
// - オークションは終了済みであること
// - 最高入札価格が設定されていること
//
// # 戻り値
// - Price 落札者が支払う金額
// - エラー
func (a *Auction) GetBuyerPrice() (*Price, error) {
	if a.highBidPrice == nil {
		return nil, NewAuctionError("high bid price is not set")
	}
	return buyerPriceCalculator.CalculatePrice(a, a.highBidPrice)
}

/**
func (a *Auction) GetBuyerPrice() (*Price, error) {
	if a.highBidPrice == nil {
		return nil, NewAuctionError("high bid price is not set")
	}
	switch a.product.GetProductType() {
	case ProductTypeGeneric:
		p, err := NewPrice(BaseShippingFee)
		if err != nil {
			return nil, err
		}
		return a.highBidPrice.Add(p), nil
	case ProductTypeDownloadSoftware:
		return a.highBidPrice, nil
	case ProductTypeCar:
		p, err := NewPrice(CarShippingFee)
		if err != nil {
			return nil, err
		}
		buyerPrice := a.highBidPrice.Add(p)
		if a.highBidPrice.IsGreaterThanOrEqualTo(&Price{LuxuryCarPriceThreshold}) {
			return buyerPrice.Add(a.highBidPrice.Multiply(LuxuryTaxRate)), nil
		} else {
			return buyerPrice, nil
		}
	default:
		return nil, NewAuctionError("unknown product type")
	}
}
*/
