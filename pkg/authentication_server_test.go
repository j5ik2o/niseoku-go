package pkg

import (
	"github.com/stretchr/testify/require"
	"nisecari-go/pkg/domain"
	"nisecari-go/pkg/infrastracture"
	"testing"
)

// 2) 登録ユーザーとして、本人確認を受けるために、ログインする
func Test_未ログインユーザがログインできる(t *testing.T) {
	userAccountRepository := infrastracture.NewUserAccountRepositoryInMemory()
	sessionRepository := infrastracture.NewSessionRepositoryInMemory()
	userAccount1 := domain.NewUserAccount(domain.GenerateUserAccountId(), "Junichi", "Kato")
	err := userAccountRepository.Store(userAccount1)
	require.NoError(t, err)
	authenticationService := infrastracture.NewAuthenticationService(userAccountRepository, sessionRepository)
	login, err := authenticationService.Login(userAccount1.Id)
	require.NoError(t, err)
	require.NotNil(t, login)
	session, err := sessionRepository.FindById(login.Session.Id)
	require.NoError(t, err)
	require.NotNil(t, session)
}

// 3) 認証されたユーザーとして、サービスの利用を終えるために、ログアウトする
func Test_ログイン済みユーザがログアウトできる(t *testing.T) {
	userAccountRepository := infrastracture.NewUserAccountRepositoryInMemory()
	sessionRepository := infrastracture.NewSessionRepositoryInMemory()
	userAccount1 := domain.NewUserAccount(domain.GenerateUserAccountId(), "Junichi", "Kato")
	err := userAccountRepository.Store(userAccount1)
	require.NoError(t, err)
	authenticationService := infrastracture.NewAuthenticationService(userAccountRepository, sessionRepository)
	login, err := authenticationService.Login(userAccount1.Id)
	require.NoError(t, err)
	require.NotNil(t, login)
	session, err := sessionRepository.FindById(login.Session.Id)
	require.NoError(t, err)
	require.NotNil(t, session)
	err = authenticationService.Logout(login.Session.Id)
	require.NoError(t, err)
	session, err = sessionRepository.FindById(login.Session.Id)
	require.NoError(t, err)
	require.Nil(t, session)
}
