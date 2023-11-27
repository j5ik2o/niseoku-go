package infrastructure

import (
	"github.com/stretchr/testify/require"
	"niseoku-go/pkg/domain"
	"niseoku-go/pkg/infrastracture/repository/memory"
	"testing"
)

// 商品を登録できる
func Test_RegisterProduct(t *testing.T) {
	// Given
	productId := domain.GenerateProductId()
	productType := domain.ProductTypeGeneric
	productName, err := domain.NewProductName("iPhone")
	require.NoError(t, err)
	productPrice, err := domain.NewProductPrice(100000)
	require.NoError(t, err)
	product := domain.NewProduct(productId, productType, productName, productPrice)
	require.NotNil(t, product)
	require.Equal(t, product.GetId(), productId)
	require.Equal(t, product.GetName(), productName)
	require.Equal(t, product.GetPrice(), productPrice)
	productRepository := memory.NewProductRepositoryInMemory()

	// When
	err = productRepository.Store(product)

	// Then
	require.NoError(t, err)
	actual, err := productRepository.FindById(product.GetId())
	require.NoError(t, err)
	require.Equal(t, product, actual)
}
