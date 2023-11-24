package domain

import "github.com/oklog/ulid/v2"

type SessionId struct {
	value string
}

func NewSessionId(value string) (*SessionId, error) {
	if value == "" {
		return nil, NewInvalidArgumentError("session id must not be empty")
	}
	return &SessionId{
		value: value,
	}, nil
}

func GenerateSessionId() *SessionId {
	ulid := ulid.Make()
	return &SessionId{
		value: ulid.String(),
	}
}

func (s *SessionId) String() string {
	return s.value
}
