package pkg

import (
	"github.com/stretchr/testify/require"
	"nisecari-go/pkg/domain"
	"testing"
	"time"
)

// 5) 認証売主として、商品を売りに出すために、オークションを作成したい。
func Test_オークションを作成できる(t *testing.T) {
	product := createProduct()
	sellerId := domain.GenerateUserAccountId()
	startDateTime := time.Now().Add(1 * time.Hour)
	endDateTime := startDateTime.Add(1 * time.Hour)
	startPrice, err := domain.NewPrice(1000)
	require.NoError(t, err)
	auction, err := domain.NewAuction(domain.GenerateAuctionId(), product, &startDateTime, &endDateTime, startPrice, sellerId)
	require.NoError(t, err)
	require.Equal(t, auction.Product, product)
	require.Equal(t, auction.SellerId, sellerId)
	require.Equal(t, auction.StartDateTime, &startDateTime)
	require.Equal(t, auction.StartPrice, startPrice)
}

func Test_開始時刻が過去の場合はエラーになる(t *testing.T) {
	product := createProduct()
	sellerId := domain.GenerateUserAccountId()
	startDateTime := time.Now().Add(-1 * time.Hour)
	endDateTime := startDateTime.Add(1 * time.Hour)
	startPrice, err := domain.NewPrice(1000)
	require.NoError(t, err)
	_, err = domain.NewAuction(domain.GenerateAuctionId(), product, &startDateTime, &endDateTime, startPrice, sellerId)
	require.Error(t, err)
}

func Test_終了時刻が開始時刻より前の場合はエラーになる(t *testing.T) {
	product := createProduct()
	sellerId := domain.GenerateUserAccountId()
	startDateTime := time.Now().Add(1 * time.Hour)
	endDateTime := startDateTime.Add(-1 * time.Hour)
	startPrice, err := domain.NewPrice(1000)
	require.NoError(t, err)
	_, err = domain.NewAuction(domain.GenerateAuctionId(), product, &startDateTime, &endDateTime, startPrice, sellerId)
	require.Error(t, err)
}

func Test_開始価格が0円の場合はエラーになる(t *testing.T) {
	_, err := domain.NewPrice(0)
	require.Error(t, err)
}

// 6) オークションとして、入札を受け付けるために、開始されたい。
func Test_オークションを開始できる(t *testing.T) {
	product := createProduct()
	sellerId := domain.GenerateUserAccountId()
	startDateTime := time.Now().Add(1 * time.Hour)
	endDateTime := startDateTime.Add(1 * time.Hour)
	startPrice, err := domain.NewPrice(1000)
	require.NoError(t, err)
	auction, err := domain.NewAuction(domain.GenerateAuctionId(), product, &startDateTime, &endDateTime, startPrice, sellerId)
	require.NoError(t, err)
	auction = auction.Start()
	require.NotNil(t, auction)
}

// 7) 認証入札者として、最高額入札者になるために、開始されたオークションに入札したい
func Test_オークションを最高額で入札する(t *testing.T) {
	// Given
	product := createProduct()
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
	auction = auction.Start()

	// When
	auction, err = auction.Bid(highBidPrice, buyerId)

	// Then
	require.NoError(t, err)
	require.Equal(t, auction.HighBidderId, buyerId)
	require.Equal(t, auction.HighBidPrice, highBidPrice)
}

func createProduct() *domain.Product {
	productId := domain.GenerateProductId()
	productName := domain.NewProductName("iPhone")
	productPrice := domain.NewProductPrice(100000)
	product := domain.NewProduct(productId, productName, productPrice)
	return product
}

func Test_オークションは最高額より少ない額で入札できない(t *testing.T) {
	// Given
	productId := domain.GenerateProductId()
	productName := domain.NewProductName("iPhone")
	productPrice := domain.NewProductPrice(100000)
	product := domain.NewProduct(productId, productName, productPrice)
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
	auction = auction.Start()

	// When
	_, err = auction.Bid(highBidPrice, buyerId)

	// Then
	require.Error(t, err)
}

// 8) オークションとして、最高入札者や売手に通知できるようになるために、閉じられたい。
func Test_オークションを終了できる_落札者が存在する場合(t *testing.T) {
	// TODO: 通知サービスを使って、落札者と売手に通知する
}

func Test_オークションを終了できる_落札者が不在の場合(t *testing.T) {
	// TODO: 通知サービスを使って、売手に通知する
}

// 9) オークションとして、手数料を扱うために、販売価格を調整したい
// - 出品者の金額は、2%の取引手数料を減算する
// - 落札者の金額は、アイテムカテゴリがダウンロードソフトウェアもしくは自動車でない限り、販売されるすべての商品に10ドルの配送料を追加する
// - 商品が自動車だったら1000ドルの配送料を追加する
// - 自動車が5万ドル以上で販売されたら、4%の贅沢税を追加する
func Test_出品者の販売価格を取得する(t *testing.T) {

}

func Test_落札者の購入価格を取得する(t *testing.T) {

}
