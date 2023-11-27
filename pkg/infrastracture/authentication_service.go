package infrastracture

import (
	"fmt"
	"niseoku-go/pkg/domain"
	"niseoku-go/pkg/infrastracture/repository"
)

// AuthenticationService は認証サービスを表します。
type AuthenticationService struct {
	userAccountRepository repository.UserAccountRepository
	sessionRepository     repository.SessionRepository
}

// NewAuthenticationService は認証サービスを生成します。
func NewAuthenticationService(userAccountRepository repository.UserAccountRepository, sessionRepository repository.SessionRepository) *AuthenticationService {
	return &AuthenticationService{
		userAccountRepository: userAccountRepository,
		sessionRepository:     sessionRepository,
	}
}

// Login はログインします。
//
// ログインする際には、以下のルールに従う必要がある
// - ユーザーアカウントが存在すること
// - すでにログインしていないこと
//
// # 引数
// - userAccountId: domain.UserAccountId
//
// # 戻り値
// - domain.SessionWithUserAccount
// - エラー
func (a *AuthenticationService) Login(userAccountId *domain.UserAccountId) (*domain.SessionWithUserAccount, error) {
	userAccount, err := a.userAccountRepository.FindById(userAccountId)
	if err != nil {
		return nil, err
	}
	if userAccount == nil {
		return nil, nil
	}
	if a.sessionRepository.ContainsByUserAccountId(userAccountId) {
		return nil, fmt.Errorf("user account (%s) is not logged in", userAccountId.String())
	}
	session := domain.NewSession(domain.GenerateSessionId(), userAccount.GetId())
	err = a.sessionRepository.Store(session)
	return &domain.SessionWithUserAccount{Session: session, UserAccount: userAccount}, nil
}

// Logout はログアウトします。
//
// ログアウトする際には、以下のルールに従う必要がある
// - ログインしていること
//
// # 引数
// - sessionId: domain.SessionId
//
// # 戻り値
// - エラー
func (a *AuthenticationService) Logout(sessionId *domain.SessionId) error {
	session, err := a.sessionRepository.FindById(sessionId)
	if err != nil {
		return err
	}
	if session == nil {
		return fmt.Errorf("session (%s) is not found", sessionId.String())
	}
	return a.sessionRepository.Delete(session)
}
