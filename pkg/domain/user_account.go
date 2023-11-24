package domain

type UserAccount struct {
	id        *UserAccountId
	firstName string
	lastName  string
}

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

func (u *UserAccount) GetId() *UserAccountId {
	return u.id
}

func (u *UserAccount) GetFirstName() string {
	return u.firstName
}

func (u *UserAccount) GetLastName() string {
	return u.lastName
}
