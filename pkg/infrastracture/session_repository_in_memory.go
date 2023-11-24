package infrastracture

import "niseoku-go/pkg/domain"

type SessionRepositoryInMemory struct {
	sessions map[*domain.SessionId]*domain.Session
}

func NewSessionRepositoryInMemory() *SessionRepositoryInMemory {
	return &SessionRepositoryInMemory{
		sessions: make(map[*domain.SessionId]*domain.Session),
	}
}

func (s *SessionRepositoryInMemory) Store(session *domain.Session) error {
	s.sessions[session.Id] = session
	return nil
}

func (s *SessionRepositoryInMemory) FindById(id *domain.SessionId) (*domain.Session, error) {
	return s.sessions[id], nil
}

func (s *SessionRepositoryInMemory) Delete(session *domain.Session) error {
	delete(s.sessions, session.Id)
	return nil
}
