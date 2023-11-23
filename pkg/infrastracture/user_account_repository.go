package infrastracture

import "nisecari-go/pkg/domain"

type UserAccountRepository interface {
	Store(userAccount *domain.UserAccount) error
	FindById(id *domain.UserAccountId) (*domain.UserAccount, error)
}
