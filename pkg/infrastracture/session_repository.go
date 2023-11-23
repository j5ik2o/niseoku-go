package infrastracture

import "nisecari-go/pkg/domain"

type SessionRepository interface {
	Store(session *domain.Session) error
	FindById(id *domain.SessionId) (*domain.Session, error)
	Delete(session *domain.Session) error
}
