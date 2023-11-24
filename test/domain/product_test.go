package domain

import (
	"github.com/stretchr/testify/require"
	"niseoku-go/pkg/domain"
	"testing"
)

func Test_商品を作成できる(t *testing.T) {
	// Given
	productId := domain.GenerateProductId()
	productType := domain.ProductTypeGeneric
	productName, err := domain.NewProductName("iPhone")
	require.NoError(t, err)
	productPrice, err := domain.NewProductPrice(100000)
	require.NoError(t, err)

	// When
	product := domain.NewProduct(productId, productType, productName, productPrice)

	// Then
	require.NotNil(t, product)
	require.Equal(t, product.Id, productId)
	require.Equal(t, product.Name, productName)
	require.Equal(t, product.Price, productPrice)
}

func Test_商品を公開できる(t *testing.T) {
	// Given
	productId := domain.GenerateProductId()
	productType := domain.ProductTypeGeneric
	productName, err := domain.NewProductName("iPhone")
	require.NoError(t, err)
	productPrice, err := domain.NewProductPrice(100000)
	require.NoError(t, err)
	product := domain.NewProduct(productId, productType, productName, productPrice)

	// When
	product = product.Publish()

	// Then
	require.NotNil(t, product)
	require.Equal(t, product.Status, domain.ProductStatusPublic)
}
