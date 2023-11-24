package domain

import (
	"github.com/stretchr/testify/require"
	"niseoku-go/pkg/domain"
	"testing"
)

func Test_ユーザアカウントを作成できる(t *testing.T) {
	// Given
	userAccountId := domain.GenerateUserAccountId()
	firstName := "Junichi"
	lastName := "Kato"

	// When
	userAccount, err := domain.NewUserAccount(userAccountId, firstName, lastName)

	// Then
	require.NoError(t, err)
	require.NotNil(t, userAccount)
	require.NotNil(t, userAccount.Id)
	require.Equal(t, "Junichi", userAccount.FirstName)
	require.Equal(t, "Kato", userAccount.LastName)
}
