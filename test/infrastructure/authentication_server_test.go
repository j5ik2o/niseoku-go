package infrastructure

import (
	"github.com/stretchr/testify/require"
	"niseoku-go/pkg/domain"
	"niseoku-go/pkg/infrastracture"
	"testing"
)

// 2) 登録ユーザーとして、本人確認を受けるために、ログインする
func Test_未ログインユーザがログインできる(t *testing.T) {
	// Given
	userAccountRepository := infrastracture.NewUserAccountRepositoryInMemory()
	sessionRepository := infrastracture.NewSessionRepositoryInMemory()
	userAccount1, err := domain.NewUserAccount(domain.GenerateUserAccountId(), "Junichi", "Kato")
	require.NoError(t, err)
	err = userAccountRepository.Store(userAccount1)
	require.NoError(t, err)
	authenticationService := infrastracture.NewAuthenticationService(userAccountRepository, sessionRepository)

	// When
	login, err := authenticationService.Login(userAccount1.Id)

	// Then
	require.NoError(t, err)
	require.NotNil(t, login)
	session, err := sessionRepository.FindById(login.Session.Id)
	require.NoError(t, err)
	require.NotNil(t, session)
}

// 3) 認証されたユーザーとして、サービスの利用を終えるために、ログアウトする
func Test_ログイン済みユーザがログアウトできる(t *testing.T) {
	// Given
	userAccountRepository := infrastracture.NewUserAccountRepositoryInMemory()
	sessionRepository := infrastracture.NewSessionRepositoryInMemory()
	userAccount1, err := domain.NewUserAccount(domain.GenerateUserAccountId(), "Junichi", "Kato")
	require.NoError(t, err)
	err = userAccountRepository.Store(userAccount1)
	require.NoError(t, err)
	authenticationService := infrastracture.NewAuthenticationService(userAccountRepository, sessionRepository)
	login, err := authenticationService.Login(userAccount1.Id)
	require.NoError(t, err)
	require.NotNil(t, login)
	session, err := sessionRepository.FindById(login.Session.Id)
	require.NoError(t, err)
	require.NotNil(t, session)

	// When
	err = authenticationService.Logout(login.Session.Id)

	// Then
	require.NoError(t, err)
	session, err = sessionRepository.FindById(login.Session.Id)
	require.NoError(t, err)
	require.Nil(t, session)
}
