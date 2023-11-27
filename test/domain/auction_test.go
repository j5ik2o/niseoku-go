package domain

import (
	"github.com/stretchr/testify/require"
	"niseoku-go/pkg/domain"
	"testing"
	"time"
)

// オークションを作成できる
func Test_CreateAuction(t *testing.T) {
	// Given
	product := createProduct(t, domain.ProductTypeGeneric)
	sellerId := domain.GenerateUserAccountId()
	startDateTime := time.Now().Add(1 * time.Hour)
	endDateTime := startDateTime.Add(1 * time.Hour)
	startPrice, err := domain.NewPrice(1000)
	require.NoError(t, err)
	auctionId := domain.GenerateAuctionId()

	// When
	auction, err := domain.NewAuction(auctionId, product, &startDateTime, &endDateTime, startPrice, sellerId)

	// Then
	require.NoError(t, err)
	require.NotNil(t, auction)
	require.Equal(t, auction.GetId(), auctionId)
	require.Equal(t, auction.GetProduct(), product)
	require.Equal(t, auction.GetSellerId(), sellerId)
	require.Equal(t, auction.GetStartDateTime(), &startDateTime)
	require.Equal(t, auction.GetStartPrice(), startPrice)
	require.Equal(t, auction.GetEndDateTime(), &endDateTime)
	require.Nil(t, auction.GetHighBidderId())
	require.Nil(t, auction.GetHighBidPrice())
	require.Nil(t, auction.GetBuyerId())
	sellerPrice, err := auction.GetSellerPrice()
	require.Error(t, err)
	require.Nil(t, sellerPrice)
	buyerPrice, err := auction.GetBuyerPrice()
	require.Error(t, err)
	require.Nil(t, buyerPrice)
}

// 開始時刻が過去の場合は、オークションは作成できない
func Test_CantCreateAuctionIfStartTimeLessThanNow(t *testing.T) {
	// Given
	product := createProduct(t, domain.ProductTypeGeneric)
	sellerId := domain.GenerateUserAccountId()
	startDateTime := time.Now().Add(-1 * time.Hour)
	endDateTime := startDateTime.Add(1 * time.Hour)
	startPrice, err := domain.NewPrice(1000)
	require.NoError(t, err)
	auctionId := domain.GenerateAuctionId()

	// When
	_, err = domain.NewAuction(auctionId, product, &startDateTime, &endDateTime, startPrice, sellerId)

	// Then
	require.Error(t, err)
}

type MockClock struct {
	now *time.Time
}

func (c *MockClock) Now() *time.Time {
	return c.now
}

// 終了時刻が開始時刻より過去の場合は、オークションは作成できない
func Test_CantCreateAuctionIfEndTimeLessThanStartTime(t *testing.T) {
	// Given
	product := createProduct(t, domain.ProductTypeGeneric)
	sellerId := domain.GenerateUserAccountId()
	startDateTime := time.Now().Add(1 * time.Hour)
	endDateTime := startDateTime.Add(-1 * time.Hour)
	startPrice, err := domain.NewPrice(1000)
	require.NoError(t, err)
	auctionId := domain.GenerateAuctionId()

	// When
	_, err = domain.NewAuction(auctionId, product, &startDateTime, &endDateTime, startPrice, sellerId)

	// Then
	require.Error(t, err)
}

// はじめて入札する
func Test_StartAuction(t *testing.T) {
	// Given
	product := createProduct(t, domain.ProductTypeGeneric)
	sellerId := domain.GenerateUserAccountId()
	startDateTime := time.Now().Add(3 * time.Second)
	endDateTime := startDateTime.Add(1 * time.Hour)
	startPrice, err := domain.NewPrice(1000)
	require.NoError(t, err)
	auctionId := domain.GenerateAuctionId()
	auction, err := domain.NewAuction(auctionId, product, &startDateTime, &endDateTime, startPrice, sellerId)
	require.NoError(t, err)
	callback := false
	setClock(startDateTime, auction)

	// When
	auction, err = auction.Start(func(auction *domain.Auction) {
		callback = true
	})

	// Then
	require.NoError(t, err)
	require.True(t, callback)
	require.NotNil(t, auction)
}

// 開始時刻前にオークションを開始できない
func Test_CantStartAuctionBeforeStartTime(t *testing.T) {
	// Given
	product := createProduct(t, domain.ProductTypeGeneric)
	sellerId := domain.GenerateUserAccountId()
	startDateTime := time.Now().Add(3 * time.Second)
	endDateTime := startDateTime.Add(1 * time.Hour)
	startPrice, err := domain.NewPrice(1000)
	require.NoError(t, err)
	auctionId := domain.GenerateAuctionId()
	auction, err := domain.NewAuction(auctionId, product, &startDateTime, &endDateTime, startPrice, sellerId)
	require.NoError(t, err)
	setClock(startDateTime.Add(-30*time.Second), auction)

	// When
	auction, err = auction.Start(func(auction *domain.Auction) {
	})

	// Then
	require.Error(t, err)
	require.Nil(t, auction)
}

// オークションが開始していない場合は、入札できない
func Test_CantBidBeforeStartTime(t *testing.T) {
	// Given
	product := createProduct(t, domain.ProductTypeGeneric)
	sellerId := domain.GenerateUserAccountId()
	startDateTime := time.Now().Add(1 * time.Hour)
	endDateTime := startDateTime.Add(1 * time.Hour)
	startPrice, err := domain.NewPrice(1000)
	require.NoError(t, err)
	auctionId := domain.GenerateAuctionId()
	auction, err := domain.NewAuction(auctionId, product, &startDateTime, &endDateTime, startPrice, sellerId)
	require.NoError(t, err)
	setClock(startDateTime, auction)

	// When
	_, err = auction.Bid(startPrice, domain.GenerateUserAccountId())

	// Then
	require.Error(t, err)
}

// 最高額にてオークションに入札する
func Test_BidHighestAmountInAuction(t *testing.T) {
	// Given
	product := createProduct(t, domain.ProductTypeGeneric)
	sellerId := domain.GenerateUserAccountId()
	startDateTime := time.Now().Add(1 * time.Hour)
	endDateTime := startDateTime.Add(1 * time.Hour)
	startPrice, err := domain.NewPrice(1000)
	require.NoError(t, err)
	auctionId := domain.GenerateAuctionId()
	auction, err := domain.NewAuction(auctionId, product, &startDateTime, &endDateTime, startPrice, sellerId)
	require.NoError(t, err)
	buyerId := domain.GenerateUserAccountId()
	highBidPrice, err := domain.NewPrice(2000)
	require.NoError(t, err)
	callback := false
	setClock(startDateTime, auction)
	auction, err = auction.Start(func(auction *domain.Auction) {
		callback = true
	})
	require.NoError(t, err)
	require.True(t, callback)

	// When
	auction, err = auction.Bid(highBidPrice, buyerId)

	// Then
	require.NoError(t, err)
	require.Equal(t, auction.GetHighBidderId(), buyerId)
	require.Equal(t, auction.GetHighBidPrice(), highBidPrice)
}

func setClock(startDateTime time.Time, auction *domain.Auction) {
	mockClock := MockClock{
		now: &startDateTime,
	}
	auction.SetClock(&mockClock)
}

// 最高額より少ない価格では入札できない
func Test_CantBidWithMinimumAmountLessThanHighestAmount(t *testing.T) {
	// Given
	product := createProduct(t, domain.ProductTypeGeneric)
	sellerId := domain.GenerateUserAccountId()
	startDateTime := time.Now().Add(1 * time.Hour)
	endDateTime := startDateTime.Add(1 * time.Hour)
	startPrice, err := domain.NewPrice(1000)
	require.NoError(t, err)
	auctionId := domain.GenerateAuctionId()
	auction, err := domain.NewAuction(auctionId, product, &startDateTime, &endDateTime, startPrice, sellerId)
	require.NoError(t, err)
	buyerId := domain.GenerateUserAccountId()
	highBidPrice, err := domain.NewPrice(10)
	require.NoError(t, err)
	callback := false
	setClock(startDateTime, auction)
	auction, err = auction.Start(func(auction *domain.Auction) {
		callback = true
	})
	require.NoError(t, err)
	require.True(t, callback)

	// When
	_, err = auction.Bid(highBidPrice, buyerId)

	// Then
	require.Error(t, err)
	require.Nil(t, auction.GetHighBidderId())
	require.Nil(t, auction.GetHighBidPrice())
}

// オークションを終了できる_落札者が存在する場合
func Test_AuctionCanBeClosed_WhenThereAreWinningBidders(t *testing.T) {
	// Given
	product := createProduct(t, domain.ProductTypeGeneric)
	sellerId := domain.GenerateUserAccountId()
	startDateTime := time.Now().Add(1 * time.Hour)
	endDateTime := startDateTime.Add(1 * time.Hour)
	startPrice, err := domain.NewPrice(1000)
	require.NoError(t, err)
	auctionId := domain.GenerateAuctionId()
	auction, err := domain.NewAuction(auctionId, product, &startDateTime, &endDateTime, startPrice, sellerId)
	require.NoError(t, err)
	buyerId := domain.GenerateUserAccountId()
	highBidPrice, err := domain.NewPrice(2000)
	require.NoError(t, err)
	callback := false
	setClock(startDateTime, auction)
	auction, err = auction.Start(func(auction *domain.Auction) {
		callback = true
	})
	require.NoError(t, err)
	require.True(t, callback)
	auction, err = auction.Bid(highBidPrice, buyerId)
	var actualBuyerId *domain.UserAccountId

	// When
	auction.Close(func(auction *domain.Auction) {
		actualBuyerId = nil
	}, func(auction *domain.Auction) {
		actualBuyerId = auction.GetBuyerId()
	})

	// Then
	require.Equal(t, actualBuyerId, buyerId)
}

// オークションを終了できる_落札者が不在の場合
func Test_AuctionCannotBeClosed_WhenThereAreNoWinningBidders(t *testing.T) {
	// Given
	product := createProduct(t, domain.ProductTypeGeneric)
	sellerId := domain.GenerateUserAccountId()
	startDateTime := time.Now().Add(1 * time.Hour)
	endDateTime := startDateTime.Add(1 * time.Hour)
	startPrice, err := domain.NewPrice(1000)
	require.NoError(t, err)
	auctionId := domain.GenerateAuctionId()
	auction, err := domain.NewAuction(auctionId, product, &startDateTime, &endDateTime, startPrice, sellerId)
	require.NoError(t, err)
	callback := false
	setClock(startDateTime, auction)
	auction, err = auction.Start(func(auction *domain.Auction) {
		callback = true
	})
	require.NoError(t, err)
	require.True(t, callback)
	var actualBuyerId *domain.UserAccountId

	// When
	auction.Close(func(auction *domain.Auction) {
		actualBuyerId = nil
	}, func(auction *domain.Auction) {
		actualBuyerId = auction.GetBuyerId()
	})

	// Then
	require.Nil(t, actualBuyerId)
}

// 9) オークションとして、手数料を扱うために、販売価格を調整したい
// - 出品者の金額は、2%の取引手数料を減算する
// - 落札者の金額は、アイテムカテゴリがダウンロードソフトウェアもしくは自動車でない限り、販売されるすべての商品に10ドルの配送料を追加する
// - 商品が自動車だったら1000ドルの配送料を追加する
// - 自動車が5万ドル以上で販売されたら、4%の贅沢税を追加する

// 出品者の販売価格を取得する_2パーセントの手数料を引く
func Test_GetSellingPrice_With2PercentCommissionDeducted(t *testing.T) {
	// Given
	product := createProduct(t, domain.ProductTypeGeneric)
	sellerId := domain.GenerateUserAccountId()
	startDateTime := time.Now().Add(1 * time.Hour)
	endDateTime := startDateTime.Add(1 * time.Hour)
	startPrice, err := domain.NewPrice(1000)
	require.NoError(t, err)
	auctionId := domain.GenerateAuctionId()
	auction, err := domain.NewAuction(auctionId, product, &startDateTime, &endDateTime, startPrice, sellerId)
	require.NoError(t, err)
	buyerId := domain.GenerateUserAccountId()
	highBidPrice, err := domain.NewPrice(2000)
	require.NoError(t, err)
	callback := false
	setClock(startDateTime, auction)
	auction, err = auction.Start(func(auction *domain.Auction) {
		callback = true
	})
	require.NoError(t, err)
	require.True(t, callback)
	auction, err = auction.Bid(highBidPrice, buyerId)

	// When
	sellerPrice, err := auction.GetSellerPrice()

	// Then
	require.NoError(t, err)
	require.Equal(t, highBidPrice.Multiply(1-0.02), sellerPrice)
}

// 落札者の購入価格を取得する_一般商品には10ドルの配送料を追加する
func Test_GetSellingPrice_WithRegularItem(t *testing.T) {
	// Given
	product := createProduct(t, domain.ProductTypeGeneric)
	sellerId := domain.GenerateUserAccountId()
	startDateTime := time.Now().Add(1 * time.Hour)
	endDateTime := startDateTime.Add(1 * time.Hour)
	startPrice, err := domain.NewPrice(1000)
	require.NoError(t, err)
	auctionId := domain.GenerateAuctionId()
	auction, err := domain.NewAuction(auctionId, product, &startDateTime, &endDateTime, startPrice, sellerId)
	require.NoError(t, err)
	buyerId := domain.GenerateUserAccountId()
	highBidPrice, err := domain.NewPrice(2000)
	require.NoError(t, err)
	callback := false
	setClock(startDateTime, auction)
	auction, err = auction.Start(func(auction *domain.Auction) {
		callback = true
	})
	require.NoError(t, err)
	require.True(t, callback)
	auction, err = auction.Bid(highBidPrice, buyerId)

	// When
	buyerPrice, err := auction.GetBuyerPrice()

	// Then
	require.NoError(t, err)
	require.Equal(t, highBidPrice.Add(domain.NewPriceFromInt(10)), buyerPrice)
}

// 落札者の購入価格を取得する_ダウンロードソフトウェア
func Test_GetSellingPrice_WithDownloadableSoftware(t *testing.T) {
	// Given
	product := createProduct(t, domain.ProductTypeDownloadSoftware)
	sellerId := domain.GenerateUserAccountId()
	startDateTime := time.Now().Add(1 * time.Hour)
	endDateTime := startDateTime.Add(1 * time.Hour)
	startPrice, err := domain.NewPrice(1000)
	require.NoError(t, err)
	auctionId := domain.GenerateAuctionId()
	auction, err := domain.NewAuction(auctionId, product, &startDateTime, &endDateTime, startPrice, sellerId)
	require.NoError(t, err)
	buyerId := domain.GenerateUserAccountId()
	highBidPrice, err := domain.NewPrice(2000)
	require.NoError(t, err)
	callback := false
	setClock(startDateTime, auction)
	auction, err = auction.Start(func(auction *domain.Auction) {
		callback = true
	})
	require.NoError(t, err)
	require.True(t, callback)
	auction, err = auction.Bid(highBidPrice, buyerId)

	// When
	buyerPrice, err := auction.GetBuyerPrice()

	// Then
	require.NoError(t, err)
	require.Equal(t, highBidPrice, buyerPrice)
}

// 落札者の購入価格を取得する_自動車(1000ドルの送料が追加)
func Test_GetSellingPrice_WithCar(t *testing.T) {
	// Given
	product := createProduct(t, domain.ProductTypeCar)
	sellerId := domain.GenerateUserAccountId()
	startDateTime := time.Now().Add(1 * time.Hour)
	endDateTime := startDateTime.Add(1 * time.Hour)
	startPrice, err := domain.NewPrice(1000)
	require.NoError(t, err)
	auctionId := domain.GenerateAuctionId()
	auction, err := domain.NewAuction(auctionId, product, &startDateTime, &endDateTime, startPrice, sellerId)
	require.NoError(t, err)
	buyerId := domain.GenerateUserAccountId()
	highBidPrice, err := domain.NewPrice(2000)
	require.NoError(t, err)
	callback := false
	setClock(startDateTime, auction)
	auction, err = auction.Start(func(auction *domain.Auction) {
		callback = true
	})
	require.NoError(t, err)
	require.True(t, callback)
	auction, err = auction.Bid(highBidPrice, buyerId)

	// When
	buyerPrice, err := auction.GetBuyerPrice()

	// Then
	require.NoError(t, err)
	require.Equal(t, highBidPrice.Add(domain.NewPriceFromInt(1000)), buyerPrice)
}

// 落札者の購入価格を取得する_5万ドル以上の自動車(4%の贅沢税追加)
func Test_GetSellingPrice_WithCarOver50K(t *testing.T) {
	// Given
	product := createProduct(t, domain.ProductTypeCar)
	sellerId := domain.GenerateUserAccountId()
	startDateTime := time.Now().Add(1 * time.Hour)
	endDateTime := startDateTime.Add(1 * time.Hour)
	startPrice, err := domain.NewPrice(1000)
	require.NoError(t, err)
	auctionId := domain.GenerateAuctionId()
	auction, err := domain.NewAuction(auctionId, product, &startDateTime, &endDateTime, startPrice, sellerId)
	require.NoError(t, err)
	buyerId := domain.GenerateUserAccountId()
	highBidPrice, err := domain.NewPrice(50000)
	require.NoError(t, err)
	callback := false
	setClock(startDateTime, auction)
	auction, err = auction.Start(func(auction *domain.Auction) {
		callback = true
	})
	require.NoError(t, err)
	require.True(t, callback)
	auction, err = auction.Bid(highBidPrice, buyerId)

	// When
	buyerPrice, err := auction.GetBuyerPrice()

	// Then
	require.NoError(t, err)
	require.Equal(t, highBidPrice.Add(domain.NewPriceFromInt(1000)).Add(highBidPrice.Multiply(0.04)), buyerPrice)
}

func createProduct(t *testing.T, productType domain.ProductType) *domain.Product {
	productId := domain.GenerateProductId()
	productName, err := domain.NewProductName("iPhone")
	require.NoError(t, err)
	productPrice, err := domain.NewProductPrice(100000)
	require.NoError(t, err)
	product := domain.NewProduct(productId, productType, productName, productPrice)
	return product
}
