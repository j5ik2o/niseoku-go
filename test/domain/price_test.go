package domain

import (
	"github.com/stretchr/testify/require"
	"niseoku-go/pkg/domain"
	"testing"
)

// 価格が0円以下の場合は価格を作成できない
func Test_CantCreatePriceIfValueLessThanZero(t *testing.T) {
	// Given, When
	_, err := domain.NewPrice(0)

	// Then
	require.Error(t, err)
}
