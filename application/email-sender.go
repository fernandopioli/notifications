package application

type EmailSender interface {
	SendEmail(email string, subject string, body string) error
}