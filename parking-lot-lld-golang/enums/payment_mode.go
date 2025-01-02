package enums

type PaymentMode int

const (
	CASH PaymentMode = iota
	CARD
	UPI
	OTHER
)

func (a PaymentMode) String() string {
	switch a {
	case CASH:
		return "Cash"
	case CARD:
		return "Card"
	case UPI:
		return "UPI"
	}
	return "Other"
}
