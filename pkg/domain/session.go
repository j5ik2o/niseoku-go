package domain

type Session struct {
	Id            *SessionId
	UserAccountId *UserAccountId
}

func NewSession(id *SessionId, userAccountId *UserAccountId) *Session {
	return &Session{
		Id:            id,
		UserAccountId: userAccountId,
	}
}
