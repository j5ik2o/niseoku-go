package infrastracture

import (
	"niseoku-go/pkg/domain"
	"niseoku-go/pkg/infrastracture/repository"
)

type AuthenticationService struct {
	userAccountRepository repository.UserAccountRepository
	sessionRepository     repository.SessionRepository
}

func NewAuthenticationService(userAccountRepository repository.UserAccountRepository, sessionRepository repository.SessionRepository) *AuthenticationService {
	return &AuthenticationService{
		userAccountRepository: userAccountRepository,
		sessionRepository:     sessionRepository,
	}
}

func (a *AuthenticationService) Login(userAccountId *domain.UserAccountId) (*domain.SessionWithUserAccount, error) {
	userAccount, err := a.userAccountRepository.FindById(userAccountId)
	if err != nil {
		return nil, err
	}
	if userAccount == nil {
		return nil, nil
	}
	session := domain.NewSession(domain.GenerateSessionId(), userAccount.Id)
	err = a.sessionRepository.Store(session)
	return &domain.SessionWithUserAccount{Session: session, UserAccount: userAccount}, nil
}

func (a *AuthenticationService) Logout(sessionId *domain.SessionId) error {
	session, err := a.sessionRepository.FindById(sessionId)
	if err != nil {
		return err
	}
	if session == nil {
		return nil
	}
	return a.sessionRepository.Delete(session)
}
