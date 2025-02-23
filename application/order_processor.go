package application

import (
	"fmt"
	"notifications/domain"
)

type OrderProcessor struct {
	emailSender EmailSender
}

func (op *OrderProcessor) ProcessOrder(order domain.Order) error {
	subject := "Pedido realizado com sucesso"
	body := fmt.Sprintf("Olá %s, seu pedido foi realizado com sucesso. O total do pedido é de R$%.2f.", order.Customer.Name, order.Total)
	return op.emailSender.SendEmail(order.Customer.Email, subject, body)
}

func NewOrderProcessor(emailSender EmailSender) *OrderProcessor {
	return &OrderProcessor{emailSender: emailSender}
}