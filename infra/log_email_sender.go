package infra

import (
	"log"
	"notifications/application"
)

type LogEmailSender struct{}

var _ application.EmailSender = (*LogEmailSender)(nil)

func (s *LogEmailSender) SendEmail(email, subject, body string) error {
	log.Printf("Enviando email para %s com assunto %s e corpo %s\n", email, subject, body)
	return nil
}


