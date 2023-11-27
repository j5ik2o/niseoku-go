package infrastructure

import "niseoku-go/pkg/infrastracture"

// EmailServiceMock はメールサービスのモックを表します。
type EmailServiceMock struct {
	notices map[string]infrastracture.Email
}

// NewEmailServiceMock はメールサービスのモックを生成します。
func NewEmailServiceMock() *EmailServiceMock {
	return &EmailServiceMock{
		notices: make(map[string]infrastracture.Email),
	}
}

// Send はメールを送信します。
func (e *EmailServiceMock) Send(to string, subject string, body string) error {
	email := infrastracture.Email{
		To:      to,
		Subject: subject,
		Body:    body,
	}
	e.notices[to] = email
	return nil
}

// Clear はメールをクリアします。
func (e *EmailServiceMock) Clear() {
	e.notices = make(map[string]infrastracture.Email)
}

// Exists はメールが存在するかどうかを判定します。
func (e *EmailServiceMock) Exists(to string, f func(email infrastracture.Email) bool) bool {
	email, ok := e.notices[to]
	if !ok {
		return false
	}
	return f(email)
}
