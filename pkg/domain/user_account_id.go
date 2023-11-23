package domain

import "github.com/oklog/ulid/v2"

type UserAccountId struct {
	Value string
}

func NewUserAccountId(value string) *UserAccountId {
	return &UserAccountId{
		Value: value,
	}
}

func GenerateUserAccountId() *UserAccountId {
	ulid := ulid.Make()
	return &UserAccountId{
		Value: ulid.String(),
	}
}

func (u *UserAccountId) Equals(other *UserAccountId) bool {
	return u.Value == other.Value
}
