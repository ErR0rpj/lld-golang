package enums

type Status int

const (
	SUCCESSFUL Status = iota
	PENDING
	UNSUCCESSFUL
)

func (a Status) String() string {
	switch a {
	case SUCCESSFUL:
		return "Successful"
	case PENDING:
		return "Pending"
	}
	return "Unsuccessful"
}
