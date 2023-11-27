package domain

import "github.com/oklog/ulid/v2"

// SessionId はセッションのIDを表します。
type SessionId struct {
	value string
}

// NewSessionId はセッションのIDを生成します。
func NewSessionId(value string) (*SessionId, error) {
	if value == "" {
		return nil, NewInvalidArgumentError("session id must not be empty")
	}
	return &SessionId{
		value: value,
	}, nil
}

// GenerateSessionId はセッションのIDを生成します。
func GenerateSessionId() *SessionId {
	ulid := ulid.Make()
	return &SessionId{
		value: ulid.String(),
	}
}

// String はセッションのIDを文字列で返します。
func (s *SessionId) String() string {
	return s.value
}
