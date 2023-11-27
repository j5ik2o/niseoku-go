package repository

import "niseoku-go/pkg/domain"

// SessionRepository はセッションのリポジトリを表します。
type SessionRepository interface {
	// Store はセッションを保存します。
	Store(session *domain.Session) error
	// FindById はセッションをIDで検索します。
	FindById(id *domain.SessionId) (*domain.Session, error)
	// ContainsByUserAccountId はユーザーアカウントIDを持つセッションが存在するかを返します。
	ContainsByUserAccountId(userAccountId *domain.UserAccountId) bool
	// Delete はセッションを削除します。
	Delete(session *domain.Session) error
}
