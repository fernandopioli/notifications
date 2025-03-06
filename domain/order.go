package domain

import "errors"

type Order struct {
	Id       string   `json:"id"`
	Total    float64  `json:"total"`
	Customer Customer `json:"customer"`
}

func NewOrder(id string, total float64, customer Customer) (*Order, error) {
	if err := ValidateOrder(id, total, customer); err != nil {
		return nil, err
	}

	return &Order{
		Id:       id,
		Total:    total,
		Customer: customer,
	}, nil
}

func ValidateOrder(id string, total float64, customer Customer) error {
	if id == "" {
		return errors.New("order id is required")
	}
	if total <= 0 {
		return errors.New("order total must be greater than 0")
	}
	if customer.Id == "" {
		return errors.New("customer is required")
	}

	return nil
}

func (o *Order) GetId() string { return o.Id }

func (o *Order) GetTotal() float64 { return o.Total }

func (o *Order) GetCustomer() *Customer { return &o.Customer }
