package domain

import (
	"github.com/stretchr/testify/require"
	"niseoku-go/pkg/domain"
	"testing"
	"time"
)

// 5) 認証売主として、商品を売りに出すために、オークションを作成したい。
func Test_オークションを作成できる(t *testing.T) {
	// Given
	product := createProduct(t, domain.ProductTypeGeneric)
	sellerId := domain.GenerateUserAccountId()
	startDateTime := time.Now().Add(1 * time.Hour)
	endDateTime := startDateTime.Add(1 * time.Hour)
	startPrice, err := domain.NewPrice(1000)
	require.NoError(t, err)

	// When
	auction, err := domain.NewAuction(domain.GenerateAuctionId(), product, &startDateTime, &endDateTime, startPrice, sellerId)

	// Then
	require.NoError(t, err)
	require.Equal(t, auction.Product, product)
	require.Equal(t, auction.SellerId, sellerId)
	require.Equal(t, auction.StartDateTime, &startDateTime)
	require.Equal(t, auction.StartPrice, startPrice)
}

func Test_開始時刻が過去の場合はエラーになる(t *testing.T) {
	// Given
	product := createProduct(t, domain.ProductTypeGeneric)
	sellerId := domain.GenerateUserAccountId()
	startDateTime := time.Now().Add(-1 * time.Hour)
	endDateTime := startDateTime.Add(1 * time.Hour)
	startPrice, err := domain.NewPrice(1000)
	require.NoError(t, err)

	// When
	_, err = domain.NewAuction(domain.GenerateAuctionId(), product, &startDateTime, &endDateTime, startPrice, sellerId)

	// Then
	require.Error(t, err)
}

func Test_終了時刻が開始時刻より前の場合はエラーになる(t *testing.T) {
	// Given
	product := createProduct(t, domain.ProductTypeGeneric)
	sellerId := domain.GenerateUserAccountId()
	startDateTime := time.Now().Add(1 * time.Hour)
	endDateTime := startDateTime.Add(-1 * time.Hour)
	startPrice, err := domain.NewPrice(1000)
	require.NoError(t, err)

	// When
	_, err = domain.NewAuction(domain.GenerateAuctionId(), product, &startDateTime, &endDateTime, startPrice, sellerId)

	// Then
	require.Error(t, err)
}

func Test_開始価格が0円の場合はエラーになる(t *testing.T) {
	// Given, When
	_, err := domain.NewPrice(0)

	// Then
	require.Error(t, err)
}

// 6) オークションとして、入札を受け付けるために、開始されたい。
func Test_オークションを開始できる(t *testing.T) {
	// Given
	product := createProduct(t, domain.ProductTypeGeneric)
	sellerId := domain.GenerateUserAccountId()
	startDateTime := time.Now().Add(1 * time.Hour)
	endDateTime := startDateTime.Add(1 * time.Hour)
	startPrice, err := domain.NewPrice(1000)
	require.NoError(t, err)
	auction, err := domain.NewAuction(domain.GenerateAuctionId(), product, &startDateTime, &endDateTime, startPrice, sellerId)
	require.NoError(t, err)
	callback := false

	// When
	auction = auction.Start(func(auction *domain.Auction) {
		callback = true
	})

	// Then
	require.True(t, callback)
	require.NotNil(t, auction)
}

// 7) 認証入札者として、最高額入札者になるために、開始されたオークションに入札したい
func Test_オークションを最高額で入札する(t *testing.T) {
	// Given
	product := createProduct(t, domain.ProductTypeGeneric)
	sellerId := domain.GenerateUserAccountId()
	startDateTime := time.Now().Add(1 * time.Hour)
	endDateTime := startDateTime.Add(1 * time.Hour)
	startPrice, err := domain.NewPrice(1000)
	require.NoError(t, err)
	auction, err := domain.NewAuction(domain.GenerateAuctionId(), product, &startDateTime, &endDateTime, startPrice, sellerId)
	require.NoError(t, err)
	buyerId := domain.GenerateUserAccountId()
	highBidPrice, err := domain.NewPrice(2000)
	require.NoError(t, err)
	callback := false
	auction = auction.Start(func(auction *domain.Auction) {
		callback = true
	})
	require.True(t, callback)

	// When
	auction, err = auction.Bid(highBidPrice, buyerId)

	// Then
	require.NoError(t, err)
	require.Equal(t, auction.HighBidderId, buyerId)
	require.Equal(t, auction.HighBidPrice, highBidPrice)
}

func Test_オークションは最高額より少ない額で入札できない(t *testing.T) {
	// Given
	product := createProduct(t, domain.ProductTypeGeneric)
	sellerId := domain.GenerateUserAccountId()
	startDateTime := time.Now().Add(1 * time.Hour)
	endDateTime := startDateTime.Add(1 * time.Hour)
	startPrice, err := domain.NewPrice(1000)
	require.NoError(t, err)
	auction, err := domain.NewAuction(domain.GenerateAuctionId(), product, &startDateTime, &endDateTime, startPrice, sellerId)
	require.NoError(t, err)
	buyerId := domain.GenerateUserAccountId()
	highBidPrice, err := domain.NewPrice(10)
	require.NoError(t, err)
	callback := false
	auction = auction.Start(func(auction *domain.Auction) {
		callback = true
	})
	require.True(t, callback)

	// When
	_, err = auction.Bid(highBidPrice, buyerId)

	// Then
	require.Error(t, err)
	require.Nil(t, auction.HighBidderId)
	require.Nil(t, auction.HighBidPrice)
}

// 8) オークションとして、最高入札者や売手に通知できるようになるために、閉じられたい。
func Test_オークションを終了できる_落札者が存在する場合(t *testing.T) {
	// Given
	product := createProduct(t, domain.ProductTypeGeneric)
	sellerId := domain.GenerateUserAccountId()
	startDateTime := time.Now().Add(1 * time.Hour)
	endDateTime := startDateTime.Add(1 * time.Hour)
	startPrice, err := domain.NewPrice(1000)
	require.NoError(t, err)
	auction, err := domain.NewAuction(domain.GenerateAuctionId(), product, &startDateTime, &endDateTime, startPrice, sellerId)
	require.NoError(t, err)
	buyerId := domain.GenerateUserAccountId()
	highBidPrice, err := domain.NewPrice(2000)
	require.NoError(t, err)
	callback := false
	auction = auction.Start(func(auction *domain.Auction) {
		callback = true
	})
	require.True(t, callback)
	auction, err = auction.Bid(highBidPrice, buyerId)
	var actualBuyerId *domain.UserAccountId

	// When
	auction.Close(func(auction *domain.Auction) {
		actualBuyerId = nil
	}, func(auction *domain.Auction) {
		actualBuyerId = auction.BuyerId
	})

	// Then
	require.Equal(t, actualBuyerId, buyerId)
}

func Test_オークションを終了できる_落札者が不在の場合(t *testing.T) {
	// Given
	product := createProduct(t, domain.ProductTypeGeneric)
	sellerId := domain.GenerateUserAccountId()
	startDateTime := time.Now().Add(1 * time.Hour)
	endDateTime := startDateTime.Add(1 * time.Hour)
	startPrice, err := domain.NewPrice(1000)
	require.NoError(t, err)
	auction, err := domain.NewAuction(domain.GenerateAuctionId(), product, &startDateTime, &endDateTime, startPrice, sellerId)
	require.NoError(t, err)
	require.NoError(t, err)
	callback := false
	auction = auction.Start(func(auction *domain.Auction) {
		callback = true
	})
	require.True(t, callback)
	var actualBuyerId *domain.UserAccountId

	// When
	auction.Close(func(auction *domain.Auction) {
		actualBuyerId = nil
	}, func(auction *domain.Auction) {
		actualBuyerId = auction.BuyerId
	})

	// Then
	require.Nil(t, actualBuyerId)
}

// 9) オークションとして、手数料を扱うために、販売価格を調整したい
// - 出品者の金額は、2%の取引手数料を減算する
// - 落札者の金額は、アイテムカテゴリがダウンロードソフトウェアもしくは自動車でない限り、販売されるすべての商品に10ドルの配送料を追加する
// - 商品が自動車だったら1000ドルの配送料を追加する
// - 自動車が5万ドル以上で販売されたら、4%の贅沢税を追加する
func Test_出品者の販売価格を取得する_2パーセントの手数料を引く(t *testing.T) {
	// Given
	product := createProduct(t, domain.ProductTypeGeneric)
	sellerId := domain.GenerateUserAccountId()
	startDateTime := time.Now().Add(1 * time.Hour)
	endDateTime := startDateTime.Add(1 * time.Hour)
	startPrice, err := domain.NewPrice(1000)
	require.NoError(t, err)
	auction, err := domain.NewAuction(domain.GenerateAuctionId(), product, &startDateTime, &endDateTime, startPrice, sellerId)
	require.NoError(t, err)
	buyerId := domain.GenerateUserAccountId()
	highBidPrice, err := domain.NewPrice(2000)
	require.NoError(t, err)
	callback := false
	auction = auction.Start(func(auction *domain.Auction) {
		callback = true
	})
	require.True(t, callback)
	auction, err = auction.Bid(highBidPrice, buyerId)

	// When
	sellerPrice, err := auction.GetSellerPrice()

	// Then
	require.NoError(t, err)
	require.Equal(t, highBidPrice.Multiply(1-0.02), sellerPrice)
}

func Test_落札者の購入価格を取得する_一般商品(t *testing.T) {
	// Given
	product := createProduct(t, domain.ProductTypeGeneric)
	sellerId := domain.GenerateUserAccountId()
	startDateTime := time.Now().Add(1 * time.Hour)
	endDateTime := startDateTime.Add(1 * time.Hour)
	startPrice, err := domain.NewPrice(1000)
	require.NoError(t, err)
	auction, err := domain.NewAuction(domain.GenerateAuctionId(), product, &startDateTime, &endDateTime, startPrice, sellerId)
	require.NoError(t, err)
	buyerId := domain.GenerateUserAccountId()
	highBidPrice, err := domain.NewPrice(2000)
	require.NoError(t, err)
	callback := false
	auction = auction.Start(func(auction *domain.Auction) {
		callback = true
	})
	require.True(t, callback)
	auction, err = auction.Bid(highBidPrice, buyerId)

	// When
	buyerPrice, err := auction.GetBuyerPrice()

	// Then
	require.NoError(t, err)
	require.Equal(t, highBidPrice.Add(&domain.Price{Value: 10}), buyerPrice)
}

func Test_落札者の購入価格を取得する_ダウンロードソフトウェア(t *testing.T) {
	// Given
	product := createProduct(t, domain.ProductTypeDownloadSoftware)
	sellerId := domain.GenerateUserAccountId()
	startDateTime := time.Now().Add(1 * time.Hour)
	endDateTime := startDateTime.Add(1 * time.Hour)
	startPrice, err := domain.NewPrice(1000)
	require.NoError(t, err)
	auction, err := domain.NewAuction(domain.GenerateAuctionId(), product, &startDateTime, &endDateTime, startPrice, sellerId)
	require.NoError(t, err)
	buyerId := domain.GenerateUserAccountId()
	highBidPrice, err := domain.NewPrice(2000)
	require.NoError(t, err)
	callback := false
	auction = auction.Start(func(auction *domain.Auction) {
		callback = true
	})
	require.True(t, callback)
	auction, err = auction.Bid(highBidPrice, buyerId)

	// When
	buyerPrice, err := auction.GetBuyerPrice()

	// Then
	require.NoError(t, err)
	require.Equal(t, highBidPrice, buyerPrice)
}

func Test_落札者の購入価格を取得する_自動車(t *testing.T) {
	// Given
	product := createProduct(t, domain.ProductTypeCar)
	sellerId := domain.GenerateUserAccountId()
	startDateTime := time.Now().Add(1 * time.Hour)
	endDateTime := startDateTime.Add(1 * time.Hour)
	startPrice, err := domain.NewPrice(1000)
	require.NoError(t, err)
	auction, err := domain.NewAuction(domain.GenerateAuctionId(), product, &startDateTime, &endDateTime, startPrice, sellerId)
	require.NoError(t, err)
	buyerId := domain.GenerateUserAccountId()
	highBidPrice, err := domain.NewPrice(2000)
	require.NoError(t, err)
	callback := false
	auction = auction.Start(func(auction *domain.Auction) {
		callback = true
	})
	require.True(t, callback)
	auction, err = auction.Bid(highBidPrice, buyerId)

	// When
	buyerPrice, err := auction.GetBuyerPrice()

	// Then
	require.NoError(t, err)
	require.Equal(t, highBidPrice.Add(&domain.Price{Value: 1000}), buyerPrice)
}

func Test_落札者の購入価格を取得する_自動車2(t *testing.T) {
	// Given
	product := createProduct(t, domain.ProductTypeCar)
	sellerId := domain.GenerateUserAccountId()
	startDateTime := time.Now().Add(1 * time.Hour)
	endDateTime := startDateTime.Add(1 * time.Hour)
	startPrice, err := domain.NewPrice(1000)
	require.NoError(t, err)
	auction, err := domain.NewAuction(domain.GenerateAuctionId(), product, &startDateTime, &endDateTime, startPrice, sellerId)
	require.NoError(t, err)
	buyerId := domain.GenerateUserAccountId()
	highBidPrice, err := domain.NewPrice(50000)
	require.NoError(t, err)
	callback := false
	auction = auction.Start(func(auction *domain.Auction) {
		callback = true
	})
	require.True(t, callback)
	auction, err = auction.Bid(highBidPrice, buyerId)

	// When
	buyerPrice, err := auction.GetBuyerPrice()

	// Then
	require.NoError(t, err)
	require.Equal(t, highBidPrice.Add(&domain.Price{Value: 1000}).Add(highBidPrice.Multiply(0.04)), buyerPrice)
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
