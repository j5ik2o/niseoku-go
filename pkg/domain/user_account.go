package domain

type UserAccount struct {
	Id        *UserAccountId
	FirstName string
	LastName  string
}

func NewUserAccount(id *UserAccountId, firstName string, lastName string) *UserAccount {
	return &UserAccount{
		Id:        id,
		FirstName: firstName,
		LastName:  lastName,
	}
}
