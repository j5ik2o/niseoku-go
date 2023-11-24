package infrastracture

type Email struct {
	To      string
	Subject string
	Body    string
}

type EmailService interface {
	Send(to string, subject string, body string) error
	Clear()
	Exists(to string, f func(email Email) bool) bool
}

type EmailServiceMock struct {
	notices map[string]Email
}

func NewEmailServiceMock() *EmailServiceMock {
	return &EmailServiceMock{
		notices: make(map[string]Email),
	}
}

func (e *EmailServiceMock) Send(to string, subject string, body string) error {
	email := Email{
		To:      to,
		Subject: subject,
		Body:    body,
	}
	e.notices[to] = email
	return nil
}

func (e *EmailServiceMock) Clear() {
	e.notices = make(map[string]Email)
}

func (e *EmailServiceMock) Exists(to string, f func(email Email) bool) bool {
	email, ok := e.notices[to]
	if !ok {
		return false
	}
	return f(email)
}
