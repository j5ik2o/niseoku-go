package domain

import "github.com/oklog/ulid/v2"

type UserAccountId struct {
	value string
}

func NewUserAccountId(value string) (*UserAccountId, error) {
	if value == "" {
		return nil, NewInvalidArgumentError("user account id must not be empty")
	}
	return &UserAccountId{
		value: value,
	}, nil
}

func GenerateUserAccountId() *UserAccountId {
	ulid := ulid.Make()
	return &UserAccountId{
		value: ulid.String(),
	}
}

func (u *UserAccountId) Equals(other *UserAccountId) bool {
	return u.value == other.value
}

func (u *UserAccountId) String() string {
	return u.value
}
