package infrastracture

// Email はメールを表します。
type Email struct {
	To      string
	Subject string
	Body    string
}

// EmailService はメールサービスを表します。
type EmailService interface {
	Send(to string, subject string, body string) error
	Clear()
	Exists(to string, f func(email Email) bool) bool
}
