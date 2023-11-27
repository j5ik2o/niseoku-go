package repository

import "niseoku-go/pkg/domain"

// SessionRepository はセッションのリポジトリを表します。
type SessionRepository interface {
	Store(session *domain.Session) error
	FindById(id *domain.SessionId) (*domain.Session, error)
	ContainsByUserAccountId(userAccountId *domain.UserAccountId) bool
	Delete(session *domain.Session) error
}
