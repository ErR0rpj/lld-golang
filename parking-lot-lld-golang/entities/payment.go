package entities

import "main/enums"

type Payment struct {
	Id            int
	PaymentStatus enums.Status
	Amount        float64
	PaymentMode   enums.PaymentMode
	Bill          *Bill
}
