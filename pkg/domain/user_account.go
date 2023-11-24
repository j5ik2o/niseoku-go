package domain

type UserAccount struct {
	Id        *UserAccountId
	FirstName string
	LastName  string
}

func NewUserAccount(id *UserAccountId, firstName string, lastName string) (*UserAccount, error) {
	if firstName == "" {
		return nil, NewInvalidArgumentError("first name must not be empty")
	}
	if lastName == "" {
		return nil, NewInvalidArgumentError("last name must not be empty")
	}
	return &UserAccount{
		Id:        id,
		FirstName: firstName,
		LastName:  lastName,
	}, nil
}
