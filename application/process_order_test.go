package application

import (
	"errors"
	"log"
	"os"
	"testing"

	"notifications/domain"
)

type MockEmailSender struct {
	CallCount  int
	ShouldFail bool
	To         string
	Subject    string
	Body       string
}

func (m *MockEmailSender) SendEmail(to, subject, body string) error {
	m.CallCount++
	m.To, m.Subject, m.Body = to, subject, body
	if m.ShouldFail {
		return errors.New("email sending failed")
	}
	return nil
}

func TestExecute_Success(t *testing.T) {
	log.SetOutput(os.Stdout)

	emailSender := &MockEmailSender{}
	uc := NewProcessOrderUseCase(emailSender)

	cust, err := domain.NewCustomer("1", "Any name", "any_email@example.com")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	err = uc.Execute("order123", 99.99, cust.GetId(), cust.GetName(), cust.GetEmail())
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if emailSender.CallCount != 1 {
		t.Errorf("expected 1 call to SendEmail, got %d", emailSender.CallCount)
	}
}

func TestExecute_EmailFailure(t *testing.T) {
	log.SetOutput(os.Stdout)

	emailSender := &MockEmailSender{ShouldFail: true}
	uc := NewProcessOrderUseCase(emailSender)

	cust, err := domain.NewCustomer("1", "Any name", "any_email@example.com")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	err = uc.Execute("order123", 99.99, cust.GetId(), cust.GetName(), cust.GetEmail())
	if err == nil || err.Error() != "email sending failed" {
		t.Errorf("expected error 'email sending failed', got %v", err)
	}
	if emailSender.CallCount != 1 {
		t.Errorf("expected 1 call to SendEmail, got %d", emailSender.CallCount)
	}
}

func TestExecute_EmailParams(t *testing.T) {
	log.SetOutput(os.Stdout)

	emailSender := &MockEmailSender{}
	uc := NewProcessOrderUseCase(emailSender)

	cust, err := domain.NewCustomer("1", "Any name", "any_email@example.com")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	err = uc.Execute("order123", 99.99, cust.GetId(), cust.GetName(), cust.GetEmail())

	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if emailSender.CallCount != 1 {
		t.Errorf("expected 1 call to SendEmail, got %d", emailSender.CallCount)
	}
	if emailSender.To != "any_email@example.com" {
		t.Errorf("expected to = any_email@example.com, got %s", emailSender.To)
	}
	if emailSender.Subject != "Order Confirmation" {
		t.Errorf("expected subject = Order Confirmation, got %s", emailSender.Subject)
	}
	if emailSender.Body != "Dear Any name, your order #order123 with total 99.99 has been processed." {
		t.Errorf("expected body = ..., got %s", emailSender.Body)
	}
}
