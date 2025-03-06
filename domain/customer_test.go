package domain

import (
	"errors"
	"testing"
)

func TestNewCustomer(t *testing.T) {
	tests := []struct {
		name          string
		id            string
		custName      string
		email         string
		expectedError error
	}{
		{
			name:          "should create a customer",
			id:            "1",
			custName:      "Any Name",
			email:         "any_email@example.com",
			expectedError: nil,
		},
		{
			name:          "should return an error if id is empty",
			id:            "",
			custName:      "Any Name",
			email:         "any_email@example.com",
			expectedError: errors.New("customer id is required"),
		},
		{
			name:          "should return an error if name is empty",
			id:            "1",
			custName:      "",
			email:         "any_email@example.com",
			expectedError: errors.New("customer name is required"),
		},
		{
			name:          "should return an error if name is whitespace",
			id:            "1",
			custName:      "   ",
			email:         "any_email@example.com",
			expectedError: errors.New("customer name is required"),
		},
		{
			name:          "should return an error if email is empty",
			id:            "1",
			custName:      "Any Name",
			email:         "",
			expectedError: errors.New("customer email is invalid"),
		},
		{
			name:          "should return an error if email is invalid",
			id:            "1",
			custName:      "Any Name",
			email:         "invalid_email",
			expectedError: errors.New("customer email is invalid"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cust, err := NewCustomer(tt.id, tt.custName, tt.email)
			if err != nil && tt.expectedError == nil || err == nil && tt.expectedError != nil || (err != nil && tt.expectedError != nil && err.Error() != tt.expectedError.Error()) {
				t.Errorf("NewCustomer() error = %v, expectedError %v", err, tt.expectedError)
				return
			}
			if err == nil && cust == nil {
				t.Error("expected non-nil customer")
			}
			if err == nil {
				if cust.GetId() != tt.id {
					t.Errorf("expected id %s, got %s", tt.id, cust.GetId())
				}
				if cust.GetName() != tt.custName {
					t.Errorf("expected name %s, got %s", tt.custName, cust.GetName())
				}
				if cust.GetEmail() != tt.email {
					t.Errorf("expected email %s, got %s", tt.email, cust.GetEmail())
				}
			}
		})
	}
}

func mustNewCustomer(t *testing.T, id, name, email string) *Customer {
	cust, err := NewCustomer(id, name, email)
	if err != nil {
		t.Fatalf("mustNewCustomer: %v", err)
	}
	return cust
}
