package entities

import (
	"main/enums"
	"time"
)

type Bill struct {
	Id         int
	Ticket     *Ticket
	ExitTime   time.Time
	Payment    *Payment
	Gate       *Gate
	BillStatus enums.Status
}

func NewBill(id int, ticket *Ticket, gate *Gate, exitTime time.Time) *Bill {
	return &Bill{
		Id:       id,
		Ticket:   ticket,
		ExitTime: exitTime,
		Gate:     gate,
	}
}
