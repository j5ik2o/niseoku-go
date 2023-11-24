package infrastructure

import (
	"github.com/stretchr/testify/require"
	"niseoku-go/pkg/domain"
	"niseoku-go/pkg/infrastracture"
	memory2 "niseoku-go/pkg/infrastracture/repository/memory"
	"testing"
)

func Test_未ログインユーザがログインできる(t *testing.T) {
	// Given
	userAccountRepository := memory2.NewUserAccountRepositoryInMemory()
	sessionRepository := memory2.NewSessionRepositoryInMemory()
	userAccount1, err := domain.NewUserAccount(domain.GenerateUserAccountId(), "Junichi", "Kato")
	require.NoError(t, err)
	err = userAccountRepository.Store(userAccount1)
	require.NoError(t, err)
	authenticationService := infrastracture.NewAuthenticationService(userAccountRepository, sessionRepository)

	// When
	login, err := authenticationService.Login(userAccount1.GetId())

	// Then
	require.NoError(t, err)
	require.NotNil(t, login)
	session, err := sessionRepository.FindById(login.Session.GetId())
	require.NoError(t, err)
	require.NotNil(t, session)
}

func Test_ログイン済みユーザがログアウトできる(t *testing.T) {
	// Given
	userAccountRepository := memory2.NewUserAccountRepositoryInMemory()
	sessionRepository := memory2.NewSessionRepositoryInMemory()
	userAccount1, err := domain.NewUserAccount(domain.GenerateUserAccountId(), "Junichi", "Kato")
	require.NoError(t, err)
	err = userAccountRepository.Store(userAccount1)
	require.NoError(t, err)
	authenticationService := infrastracture.NewAuthenticationService(userAccountRepository, sessionRepository)
	login, err := authenticationService.Login(userAccount1.GetId())
	require.NoError(t, err)
	require.NotNil(t, login)
	session, err := sessionRepository.FindById(login.Session.GetId())
	require.NoError(t, err)
	require.NotNil(t, session)

	// When
	err = authenticationService.Logout(login.Session.GetId())

	// Then
	require.NoError(t, err)
	session, err = sessionRepository.FindById(login.Session.GetId())
	require.NoError(t, err)
	require.Nil(t, session)
}
