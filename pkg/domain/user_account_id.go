package domain

import "github.com/oklog/ulid/v2"

type UserAccountId struct {
	Value string
}

func NewUserAccountId(value string) (*UserAccountId, error) {
	if value == "" {
		return nil, NewInvalidArgumentError("user account id must not be empty")
	}
	return &UserAccountId{
		Value: value,
	}, nil
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
