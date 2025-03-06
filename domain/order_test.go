package domain

import (
	"errors"
	"testing"
)

func TestNewOrder(t *testing.T) {
	validCustomer := mustNewCustomer(t, "1", "Any name", "any_email@example.com")
	tests := []struct {
		name          string
		id            string
		total         float64
		customer      Customer
		expectedError error
	}{
		{
			name:          "should create an order",
			id:            "1",
			total:         100.0,
			customer:      *validCustomer,
			expectedError: nil,
		},
		{
			name:          "should return an error if id is empty",
			id:            "",
			total:         100.0,
			customer:      *validCustomer,
			expectedError: errors.New("order id is required"),
		},
		{
			name:          "should return an error if total is negative",
			id:            "1",
			total:         -100.0,
			customer:      *validCustomer,
			expectedError: errors.New("order total must be greater than 0"),
		},
		{
			name:          "should return an error if total is zero",
			id:            "1",
			total:         0.0,
			customer:      *validCustomer,
			expectedError: errors.New("order total must be greater than 0"),
		},
		{
			name:          "should return an error if customer is nil",
			id:            "1",
			total:         100.0,
			customer:      Customer{},
			expectedError: errors.New("customer is required"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			order, err := NewOrder(tt.id, tt.total, tt.customer)
			if err != nil && tt.expectedError == nil || err == nil && tt.expectedError != nil || (err != nil && tt.expectedError != nil && err.Error() != tt.expectedError.Error()) {
				t.Errorf("NewOrder() error = %v, expectedError %v", err, tt.expectedError)
				return
			}
			if err == nil && order == nil {
				t.Error("expected non-nil order")
			}
			if err == nil {
				if order.GetId() != tt.id {
					t.Errorf("expected id %s, got %s", tt.id, order.GetId())
				}
				if order.GetTotal() != tt.total {
					t.Errorf("expected total %f, got %f", tt.total, order.GetTotal())
				}
				if order.GetCustomer().GetId() != tt.customer.GetId() {
					t.Errorf("expected customer id %s, got %s", tt.customer.GetId(), order.GetCustomer().GetId())
				}
			}
		})
	}
}
