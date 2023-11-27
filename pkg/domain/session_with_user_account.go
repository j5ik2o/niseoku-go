package domain

// SessionWithUserAccount はセッションとユーザーアカウントを表します。
type SessionWithUserAccount struct {
	Session     *Session
	UserAccount *UserAccount
}
