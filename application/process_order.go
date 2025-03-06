package application

import (
	"fmt"
	"log"
	"notifications/domain"
)

type ProcessOrderUseCase struct {
	emailSender EmailSender
}

func NewProcessOrderUseCase(emailSender EmailSender) *ProcessOrderUseCase {
	return &ProcessOrderUseCase{emailSender: emailSender}
}

func (op *ProcessOrderUseCase) Execute(id string, total float64, customerId string, customerName string, customerEmail string) error {

	customer, err := domain.NewCustomer(customerId, customerName, customerEmail)
	if err != nil {
		return err
	}

	order, err := domain.NewOrder(id, total, *customer)
	if err != nil {
		return err
	}

	subject := "Order Confirmation"
	body := fmt.Sprintf("Dear %s, your order #%s with total %.2f has been processed.",
		order.GetCustomer().GetName(), order.GetId(), order.GetTotal())

	err = op.emailSender.SendEmail(order.GetCustomer().GetEmail(), subject, body)
	if err != nil {
		return err
	}

	log.Printf("Order %s processed successfully for %s", order.GetId(), customer.GetName())
	return nil
}
