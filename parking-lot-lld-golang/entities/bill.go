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
