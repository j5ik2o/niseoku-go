package infrastructure

import (
	"niseoku-go/pkg/domain"
	"niseoku-go/pkg/infrastracture/repository/memory"
	"testing"
)
import "github.com/stretchr/testify/require"

// 1) ユーザーとして、サービスを利用できるようになるために、アカウントを登録する
func Test_ユーザアカウントが保存できる(t *testing.T) {
	// Given
	userAccount, err := domain.NewUserAccount(domain.GenerateUserAccountId(), "Junichi", "Kato")
	require.NoError(t, err)
	require.NotNil(t, userAccount)
	require.NotNil(t, userAccount.Id)
	require.Equal(t, "Junichi", userAccount.FirstName)
	require.Equal(t, "Kato", userAccount.LastName)
	userAccountRepository := memory.NewUserAccountRepositoryInMemory()

	// When
	err = userAccountRepository.Store(userAccount)

	// Then
	require.NoError(t, err)
	actual, err := userAccountRepository.FindById(userAccount.Id)
	require.NoError(t, err)
	require.Equal(t, userAccount, actual)
}
