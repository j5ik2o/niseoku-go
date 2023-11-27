package repository

import "niseoku-go/pkg/domain"

// UserAccountRepository はユーザーアカウントのリポジトリを表します。
type UserAccountRepository interface {
	Store(userAccount *domain.UserAccount) error
	FindById(id *domain.UserAccountId) (*domain.UserAccount, error)
}
