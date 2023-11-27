package domain

// Session はセッションを表します。
type Session struct {
	id            *SessionId
	userAccountId *UserAccountId
}

// NewSession はセッションを生成します。
func NewSession(id *SessionId, userAccountId *UserAccountId) *Session {
	return &Session{
		id:            id,
		userAccountId: userAccountId,
	}
}

// GetId はセッションのIDを表します。
func (s *Session) GetId() *SessionId {
	return s.id
}

// GetUserAccountId はユーザーのIDを表します。
func (s *Session) GetUserAccountId() *UserAccountId {
	return s.userAccountId
}
