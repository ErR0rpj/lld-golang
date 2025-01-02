package enums

type Availability int

const (
	AVAILABLE Availability = iota
	NOT_AVAILABLE
	DISABLED
)

func (a Availability) String() string {
	switch a {
	case AVAILABLE:
		return "Available"
	case NOT_AVAILABLE:
		return "Not Available"
	case DISABLED:
		return "Disabled"
	}
	return ""
}
