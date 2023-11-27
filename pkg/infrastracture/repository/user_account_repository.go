package repository

import "niseoku-go/pkg/domain"

// UserAccountRepository はユーザーアカウントのリポジトリを表します。
type UserAccountRepository interface {
	// Store はユーザーアカウントを保存します。
	Store(userAccount *domain.UserAccount) error
	// FindById はユーザーアカウントをIDで検索します。
	FindById(id *domain.UserAccountId) (*domain.UserAccount, error)
	// FindByFullName はユーザーアカウントをフルネームで検索します。
	FindByFullName(fullName string) (*domain.UserAccount, error)
}
