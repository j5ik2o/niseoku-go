package infrastructure

import (
	"niseoku-go/pkg/domain"
	"niseoku-go/pkg/infrastracture/repository/memory"
	"testing"
)
import "github.com/stretchr/testify/require"

func Test_ユーザアカウントを登録できる(t *testing.T) {
	// Given
	userAccount, err := domain.NewUserAccount(domain.GenerateUserAccountId(), "Junichi", "Kato")
	require.NoError(t, err)
	require.NotNil(t, userAccount)
	require.NotNil(t, userAccount.GetId())
	require.Equal(t, "Junichi", userAccount.GetFirstName())
	require.Equal(t, "Kato", userAccount.GetLastName())
	userAccountRepository := memory.NewUserAccountRepositoryInMemory()

	// When
	err = userAccountRepository.Store(userAccount)

	// Then
	require.NoError(t, err)
	actual, err := userAccountRepository.FindById(userAccount.GetId())
	require.NoError(t, err)
	require.Equal(t, userAccount, actual)
}
