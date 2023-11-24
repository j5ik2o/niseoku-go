package domain

import (
	"github.com/stretchr/testify/require"
	"niseoku-go/pkg/domain"
	"testing"
)

// 4) 認証されたユーザーとして、商品を出品するために、商品を登録する
func Test_商品を作成できる(t *testing.T) {
	productId := domain.GenerateProductId()
	productType := domain.ProductTypeGeneric
	productName, err := domain.NewProductName("iPhone")
	require.NoError(t, err)
	productPrice, err := domain.NewProductPrice(100000)
	require.NoError(t, err)
	product := domain.NewProduct(productId, productType, productName, productPrice)
	require.NotNil(t, product)
	require.Equal(t, product.Id, productId)
	require.Equal(t, product.Name, productName)
	require.Equal(t, product.Price, productPrice)
}

// 5) 商品として、注文を受け付けるために公開する
func Test_商品を公開できる(t *testing.T) {
	productId := domain.GenerateProductId()
	productType := domain.ProductTypeGeneric
	productName, err := domain.NewProductName("iPhone")
	require.NoError(t, err)
	productPrice, err := domain.NewProductPrice(100000)
	require.NoError(t, err)
	product := domain.NewProduct(productId, productType, productName, productPrice)

	product = product.Publish()
	require.NotNil(t, product)
	require.Equal(t, product.Status, domain.ProductStatusPublic)
}
