package domain

import (
	"github.com/stretchr/testify/require"
	"niseoku-go/pkg/domain"
	"testing"
)

func Test_ユーザアカウントを作成できる(t *testing.T) {
	userAccount := domain.NewUserAccount(domain.GenerateUserAccountId(), "Junichi", "Kato")
	require.NotNil(t, userAccount)
	require.NotNil(t, userAccount.Id)
	require.Equal(t, "Junichi", userAccount.FirstName)
	require.Equal(t, "Kato", userAccount.LastName)
}
