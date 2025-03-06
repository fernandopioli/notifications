package infra

import (
	"bytes"
	"log"
	"testing"
)

func TestNewLoggerEmailSender(t *testing.T) {
	sender := NewLoggerEmailSender()
	if sender == nil {
		t.Fatal("expected non-nil LogEmailSender")
	}
	if sender.logger == nil {
		t.Fatal("expected non-nil logger")
	}
}

func TestLogEmailSender_SendEmail(t *testing.T) {
	var buf bytes.Buffer
	logger := log.New(&buf, "EMAIL: ", log.LstdFlags)
	sender := &LogEmailSender{logger: logger}

	err := sender.SendEmail("any_email@example.com", "Test Subject", "Test Body")
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	logOutput := buf.String()
	expected := "EMAIL: "
	if !bytes.Contains([]byte(logOutput), []byte(expected)) {
		t.Errorf("expected log to contain %q, got %q", expected, logOutput)
	}
}
