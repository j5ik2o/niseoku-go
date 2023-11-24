package domain

import "github.com/oklog/ulid/v2"

type SessionId struct {
	Value string
}

func NewSessionId(value string) (*SessionId, error) {
	if value == "" {
		return nil, NewInvalidArgumentError("session id must not be empty")
	}
	return &SessionId{
		Value: value,
	}, nil
}

func GenerateSessionId() *SessionId {
	ulid := ulid.Make()
	return &SessionId{
		Value: ulid.String(),
	}
}
