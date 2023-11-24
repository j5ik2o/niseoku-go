package domain

type Session struct {
	id            *SessionId
	userAccountId *UserAccountId
}

func NewSession(id *SessionId, userAccountId *UserAccountId) *Session {
	return &Session{
		id:            id,
		userAccountId: userAccountId,
	}
}

func (s *Session) GetId() *SessionId {
	return s.id
}

func (s *Session) GetUserAccountId() *UserAccountId {
	return s.userAccountId
}
