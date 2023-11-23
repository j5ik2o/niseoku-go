package domain

import "github.com/oklog/ulid/v2"

type SessionId struct {
	Value string
}

func NewSessionId(value string) *SessionId {
	return &SessionId{
		Value: value,
	}
}

func GenerateSessionId() *SessionId {
	ulid := ulid.Make()
	return &SessionId{
		Value: ulid.String(),
	}
}
