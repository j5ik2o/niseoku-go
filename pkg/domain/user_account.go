package domain

// UserAccount はユーザーアカウントを表します。
type UserAccount struct {
	id        *UserAccountId
	firstName string
	lastName  string
	password  string
}

// NewUserAccount はユーザーアカウントを生成します。
func NewUserAccount(id *UserAccountId, firstName string, lastName string) (*UserAccount, error) {
	if firstName == "" {
		return nil, NewInvalidArgumentError("first name must not be empty")
	}
	if lastName == "" {
		return nil, NewInvalidArgumentError("last name must not be empty")
	}
	return &UserAccount{
		id:        id,
		firstName: firstName,
		lastName:  lastName,
	}, nil
}

// GetId はユーザーアカウントのIDを返します。
func (u *UserAccount) GetId() *UserAccountId {
	return u.id
}

// GetFirstName はユーザーアカウントの名を返します。
func (u *UserAccount) GetFirstName() string {
	return u.firstName
}

// GetLastName はユーザーアカウントの姓を返します。
func (u *UserAccount) GetLastName() string {
	return u.lastName
}

func (u *UserAccount) IsPasswordCorrect(password string) bool {
	return u.password == password
}

func (u *UserAccount) GetFullName() string {
	return u.firstName + " " + u.lastName
}

func (u *UserAccount) GetPassword() string {
	return u.password
}
