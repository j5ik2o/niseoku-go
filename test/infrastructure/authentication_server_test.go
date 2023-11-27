package infrastructure

import (
	"github.com/stretchr/testify/require"
	"niseoku-go/pkg/domain"
	"niseoku-go/pkg/infrastracture"
	"niseoku-go/pkg/infrastracture/repository/memory"
	"testing"
)

// 登録済みの未ログインユーザがログインできる
func Test_RegisteredNonLoggedInUserCanLogin(t *testing.T) {
	// Given
	userAccountRepository := memory.NewUserAccountRepositoryInMemory()
	sessionRepository := memory.NewSessionRepositoryInMemory()
	userAccount1, err := domain.NewUserAccount(domain.GenerateUserAccountId(), "Junichi", "Kato")
	require.NoError(t, err)
	err = userAccountRepository.Store(userAccount1)
	require.NoError(t, err)
	authenticationService := infrastracture.NewAuthenticationService(userAccountRepository, sessionRepository)

	// When
	login, err := authenticationService.Login(userAccount1.GetId(), userAccount1.GetPassword())

	// Then
	require.NoError(t, err)
	require.NotNil(t, login)
	session, err := sessionRepository.FindById(login.Session.GetId())
	require.NoError(t, err)
	require.NotNil(t, session)
}

// 登録されていないユーザはログインできない
func Test_NonRegisteredUserCantLogin(t *testing.T) {
	// Given
	userAccountRepository := memory.NewUserAccountRepositoryInMemory()
	sessionRepository := memory.NewSessionRepositoryInMemory()
	userAccount1, err := domain.NewUserAccount(domain.GenerateUserAccountId(), "Junichi", "Kato")
	require.NoError(t, err)
	authenticationService := infrastracture.NewAuthenticationService(userAccountRepository, sessionRepository)

	// When
	login, err := authenticationService.Login(userAccount1.GetId(), userAccount1.GetPassword())

	// Then
	require.NoError(t, err)
	require.Nil(t, login)
}

// 登録済みユーザがパスワードを間違うとログインできない
func Test_RegisteredUserCantLoginIfPasswordIsWrong(t *testing.T) {
	// Given
	userAccountRepository := memory.NewUserAccountRepositoryInMemory()
	sessionRepository := memory.NewSessionRepositoryInMemory()
	userAccount1, err := domain.NewUserAccount(domain.GenerateUserAccountId(), "Junichi", "Kato")
	require.NoError(t, err)
	err = userAccountRepository.Store(userAccount1)
	require.NoError(t, err)
	authenticationService := infrastracture.NewAuthenticationService(userAccountRepository, sessionRepository)

	// When
	login, err := authenticationService.Login(userAccount1.GetId(), "wrong password")

	// Then
	require.Error(t, err)
	require.Nil(t, login)
}

// 登録済みのログイン済みユーザがログアウトできる
func Test_RegisteredLoggedInUserCanLogout(t *testing.T) {
	// Given
	userAccountRepository := memory.NewUserAccountRepositoryInMemory()
	sessionRepository := memory.NewSessionRepositoryInMemory()
	userAccount1, err := domain.NewUserAccount(domain.GenerateUserAccountId(), "Junichi", "Kato")
	require.NoError(t, err)
	err = userAccountRepository.Store(userAccount1)
	require.NoError(t, err)
	authenticationService := infrastracture.NewAuthenticationService(userAccountRepository, sessionRepository)
	login, err := authenticationService.Login(userAccount1.GetId(), userAccount1.GetPassword())
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
