package infrastructure

import (
	"github.com/stretchr/testify/require"
	"niseoku-go/pkg/domain"
	"niseoku-go/pkg/infrastracture"
	"testing"
)

// 4) 認証されたユーザーとして、商品を出品するために、商品を登録する
func Test_商品を登録できる(t *testing.T) {
	productId := domain.GenerateProductId()
	productName, err := domain.NewProductName("iPhone")
	require.NoError(t, err)
	productPrice, err := domain.NewProductPrice(100000)
	require.NoError(t, err)
	product := domain.NewProduct(productId, productName, productPrice)
	require.NotNil(t, product)
	require.Equal(t, product.Id, productId)
	require.Equal(t, product.Name, productName)
	require.Equal(t, product.Price, productPrice)

	productRepository := infrastracture.NewProductRepositoryInMemory()
	err = productRepository.Store(product)
	require.NoError(t, err)
	actual, err := productRepository.FindById(product.Id)
	require.NoError(t, err)
	require.Equal(t, product, actual)
}
