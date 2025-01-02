package entities

import "main/enums"

type Payment struct {
	Id            int
	PaymentStatus enums.Status
	Amount        float64
	PaymentMode   enums.PaymentMode
	Bill          *Bill
}

func NewPayment(id int, amount float64, bill *Bill) *Payment {
	return &Payment{
		Id:     id,
		Amount: amount,
		Bill:   bill,
	}
}
