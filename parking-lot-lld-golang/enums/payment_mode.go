package enums

type PaymentMode int

const (
	CASH PaymentMode = iota
	CARD
	UPI
	OTHER
)
