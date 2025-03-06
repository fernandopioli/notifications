package infra

import (
	"log"
	"notifications/application"
	"os"
)

type LogEmailSender struct {
	logger *log.Logger
}

var _ application.EmailSender = (*LogEmailSender)(nil)

func NewLoggerEmailSender() *LogEmailSender {
	return &LogEmailSender{
		logger: log.New(os.Stdout, "EMAIL: ", log.LstdFlags),
	}
}

func (s *LogEmailSender) SendEmail(email, subject, body string) error {
	s.logger.Printf("Sending email to %s - Subject: %s, Body: %s\n", email, subject, body)
	return nil
}
