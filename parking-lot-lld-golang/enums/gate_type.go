package enums

type GateType int

const (
	BOTH GateType = iota
	ENTRY
	EXIT
)

func (a GateType) String() string {
	switch a {
	case ENTRY:
		return "Entry"
	case EXIT:
		return "Exit"
	case BOTH:
		return "Both"
	}
	return ""
}
