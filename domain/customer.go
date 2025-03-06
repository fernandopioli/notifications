package domain

import (
	"errors"
	"net/mail"
	"strings"
)

type Customer struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func NewCustomer(id, name, email string) (*Customer, error) {
	if err := ValidateCustomer(id, name, email); err != nil {
		return nil, err
	}

	return &Customer{
		Id:    id,
		Name:  name,
		Email: email,
	}, nil
}

func ValidateCustomer(id, name, email string) error {
	if id == "" {
		return errors.New("customer id is required")
	}
	if strings.TrimSpace(name) == "" {
		return errors.New("customer name is required")
	}
	if _, err := mail.ParseAddress(email); err != nil {
		return errors.New("customer email is invalid")
	}

	return nil
}

func (c *Customer) GetId() string { return c.Id }

func (c *Customer) GetName() string { return c.Name }

func (c *Customer) GetEmail() string { return c.Email }
