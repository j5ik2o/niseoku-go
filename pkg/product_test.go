package pkg

import (
	"github.com/stretchr/testify/require"
	"nisecari-go/pkg/domain"
	"nisecari-go/pkg/infrastracture"
	"testing"
)

// 4) 認証されたユーザーとして、商品を出品するために、商品を登録する
func Test_商品を作成できる(t *testing.T) {
	productId := domain.GenerateProductId()
	productName := domain.NewProductName("iPhone")
	productPrice := domain.NewProductPrice(100000)
	product := domain.NewProduct(productId, productName, productPrice)
	require.NotNil(t, product)
	require.Equal(t, product.Id, productId)
	require.Equal(t, product.Name, productName)
	require.Equal(t, product.Price, productPrice)
}

func Test_商品を登録できる(t *testing.T) {
	productId := domain.GenerateProductId()
	productName := domain.NewProductName("iPhone")
	productPrice := domain.NewProductPrice(100000)
	product := domain.NewProduct(productId, productName, productPrice)
	require.NotNil(t, product)
	require.Equal(t, product.Id, productId)
	require.Equal(t, product.Name, productName)
	require.Equal(t, product.Price, productPrice)

	productRepository := infrastracture.NewProductRepositoryInMemory()
	err := productRepository.Store(product)
	require.NoError(t, err)
	actual, err := productRepository.FindById(product.Id)
	require.NoError(t, err)
	require.Equal(t, product, actual)
}

// 5) 商品として、注文を受け付けるために公開する
func Test_商品を公開できる(t *testing.T) {
	productId := domain.GenerateProductId()
	productName := domain.NewProductName("iPhone")
	productPrice := domain.NewProductPrice(100000)
	product := domain.NewProduct(productId, productName, productPrice)

	product = product.Publish()
	require.NotNil(t, product)
	require.Equal(t, product.Status, domain.ProductStatusPublic)
}
