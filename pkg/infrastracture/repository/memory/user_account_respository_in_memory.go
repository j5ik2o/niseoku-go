package memory

import "niseoku-go/pkg/domain"

type UserAccountRepositoryInMemory struct {
	storage map[*domain.UserAccountId]*domain.UserAccount
}

func NewUserAccountRepositoryInMemory() *UserAccountRepositoryInMemory {
	return &UserAccountRepositoryInMemory{
		storage: make(map[*domain.UserAccountId]*domain.UserAccount),
	}
}

func (u *UserAccountRepositoryInMemory) Store(userAccount *domain.UserAccount) error {
	u.storage[userAccount.GetId()] = userAccount
	return nil
}

func (u *UserAccountRepositoryInMemory) FindById(id *domain.UserAccountId) (*domain.UserAccount, error) {
	return u.storage[id], nil
}

func (u *UserAccountRepositoryInMemory) Delete(userAccount *domain.UserAccount) error {
	delete(u.storage, userAccount.GetId())
	return nil
}
