package infrastructure

import (
	"niseoku-go/pkg/domain"
	"niseoku-go/pkg/infrastracture/repository/memory"
	"testing"
)
import "github.com/stretchr/testify/require"

// ユーザアカウントを登録できる
func Test_RegisterUserAccount(t *testing.T) {
	// Given
	userAccountRepository := memory.NewUserAccountRepositoryInMemory()
	userAccount, err := createUserAccount(t)
	require.NoError(t, err)

	// When
	err = userAccountRepository.Store(userAccount)

	// Then
	require.NoError(t, err)
	actual, err := userAccountRepository.FindById(userAccount.GetId())
	require.NoError(t, err)
	require.Equal(t, userAccount, actual)
}

// ユーザアカウントを検索する(ID)
func Test_FindById(t *testing.T) {
	// Given
	userAccountRepository := memory.NewUserAccountRepositoryInMemory()
	userAccount, err := createUserAccount(t)
	require.NoError(t, err)
	err = userAccountRepository.Store(userAccount)
	require.NoError(t, err)

	// When
	actualUserAccount, err := userAccountRepository.FindById(userAccount.GetId())
	require.NoError(t, err)
	require.Equal(t, userAccount, actualUserAccount)
}

// ユーザアカウントを検索する(FullName)
func Test_FindByFullName(t *testing.T) {
	// Given
	userAccountRepository := memory.NewUserAccountRepositoryInMemory()
	userAccount, err := createUserAccount(t)
	require.NoError(t, err)
	err = userAccountRepository.Store(userAccount)
	require.NoError(t, err)

	// When
	actualUserAccount, err := userAccountRepository.FindByFullName(userAccount.GetFullName())
	require.NoError(t, err)
	require.Equal(t, userAccount, actualUserAccount)
}

func createUserAccount(t *testing.T) (*domain.UserAccount, error) {
	userAccount, err := domain.NewUserAccount(domain.GenerateUserAccountId(), "Junichi", "Kato")
	require.NoError(t, err)
	require.NotNil(t, userAccount)
	require.NotNil(t, userAccount.GetId())
	require.Equal(t, "Junichi", userAccount.GetFirstName())
	require.Equal(t, "Kato", userAccount.GetLastName())
	return userAccount, err
}
