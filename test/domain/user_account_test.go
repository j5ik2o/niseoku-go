package domain

import (
	"github.com/stretchr/testify/require"
	"niseoku-go/pkg/domain"
	"testing"
)

// ユーザアカウントを作成できる
func Test_CreateUserAccount(t *testing.T) {
	// Given
	userAccountId := domain.GenerateUserAccountId()
	firstName := "Junichi"
	lastName := "Kato"

	// When
	userAccount, err := domain.NewUserAccount(userAccountId, firstName, lastName)

	// Then
	require.NoError(t, err)
	require.NotNil(t, userAccount)
	require.NotNil(t, userAccount.GetId())
	require.Equal(t, "Junichi", userAccount.GetFirstName())
	require.Equal(t, "Kato", userAccount.GetLastName())
}
