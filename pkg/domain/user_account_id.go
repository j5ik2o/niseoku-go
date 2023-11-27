package domain

import "github.com/oklog/ulid/v2"

// UserAccountId はユーザーアカウントのIDを表します。
type UserAccountId struct {
	value string
}

// NewUserAccountId はユーザーアカウントのIDを生成します。
func NewUserAccountId(value string) (*UserAccountId, error) {
	if value == "" {
		return nil, NewInvalidArgumentError("user account id must not be empty")
	}
	return &UserAccountId{
		value: value,
	}, nil
}

// GenerateUserAccountId はユーザーアカウントのIDを生成します。
func GenerateUserAccountId() *UserAccountId {
	ulid := ulid.Make()
	return &UserAccountId{
		value: ulid.String(),
	}
}

// Equals はユーザーアカウントのIDが等しいかどうかを返します。
func (u *UserAccountId) Equals(other *UserAccountId) bool {
	return u.value == other.value
}

// String はユーザーアカウントのIDを文字列で返します。
func (u *UserAccountId) String() string {
	return u.value
}
