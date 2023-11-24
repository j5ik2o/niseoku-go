package infrastructure

import (
	"niseoku-go/pkg/domain"
	"niseoku-go/pkg/infrastracture"
	"testing"
)
import "github.com/stretchr/testify/require"

// 1) ユーザーとして、サービスを利用できるようになるために、アカウントを登録する
func Test_ユーザアカウントが保存できる(t *testing.T) {
	userAccount := domain.NewUserAccount(domain.GenerateUserAccountId(), "Junichi", "Kato")
	require.NotNil(t, userAccount)
	require.NotNil(t, userAccount.Id)
	require.Equal(t, "Junichi", userAccount.FirstName)
	require.Equal(t, "Kato", userAccount.LastName)
	userAccountRepository := infrastracture.NewUserAccountRepositoryInMemory()
	err := userAccountRepository.Store(userAccount)
	require.NoError(t, err)
	actual, err := userAccountRepository.FindById(userAccount.Id)
	require.NoError(t, err)
	require.Equal(t, userAccount, actual)
}
